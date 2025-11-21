package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"

	jwtLib "dm_loanservice/drivers/jwt"
	"dm_loanservice/drivers/logger"
	"dm_loanservice/drivers/models"
	"dm_loanservice/drivers/utils"
	ctxDM "dm_loanservice/drivers/utils/context"
)

const SessionID = "Session_Id"

func httpError(w http.ResponseWriter, statusCode int, err string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(err)
}

func httpSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func AuthMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			data := ctx.Value(ctxDM.AppSession)
			ctxSess := data.(*ctxDM.Context)
			header := ctxSess.Header.(http.Header)["Authorization"]
			var authHeader string
			if len(header) > 0 {
				authHeader = header[0]
			}
			bearerToken := strings.Split(authHeader, " ")

			if len(bearerToken) == 2 {
				userSession, errExtract := jwtLib.NewExtractTokenMetadata(authHeader)
				if errExtract != nil {
					return nil, errExtract
				}

				ctxSess.UserSession = userSession

				return next(ctx, request)
			}

			return nil, errors.New("invalid token")
		}
	}
}

func TokenVerify(ctx context.Context, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			// Parsing the access token metadata
			token, err := jwtLib.ParseAccessToken(authHeader)
			if err != nil {
				httpError(w, http.StatusUnauthorized, err.Error())
				return
			}

			if token.Valid {
				accessUuid, emailExtract, errExtract := jwtLib.ExtractTokenMetadata(authHeader)
				if errExtract != nil {
					httpError(w, http.StatusUnauthorized, errExtract.Error())
					return
				}

				emailAuth, errAuth := jwtLib.FetchAuth(ctx, accessUuid)
				if errAuth != nil {
					httpError(w, http.StatusUnauthorized, errAuth.Error())
					return
				}

				if emailExtract == emailAuth {
					next.ServeHTTP(w, r)
				} else {
					httpError(w, http.StatusUnauthorized, "Invalid Authentication")
					return
				}
			} else {
				httpError(w, http.StatusUnauthorized, err.Error())
				return
			}
		} else {
			httpError(w, http.StatusUnauthorized, "Invalid token")
			return
		}
	})
}

type Middleware func(endpoint.Endpoint) endpoint.Endpoint

func SetSession() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			reqId := utils.GenerateThreadId()
			ctx = context.WithValue(ctx, SessionID, reqId)
			return next(ctx, request)
		}
	}
}

func AuthorizationMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			resp := utils.IsAuthorized(ctx, request)
			if resp == false {
				//	TODO: return unauthorized
			}
			return next(ctx, request)
		}
	}
}

func MfaMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			data := ctx.Value(ctxDM.AppSession)
			ctxSess, ok := data.(*ctxDM.Context)
			if !ok {
				return nil, errors.New("invalid context session")
			}

			// Extract Authorization header
			header, ok := ctxSess.Header.(http.Header)
			if !ok {
				return nil, errors.New("invalid header format")
			}

			authHeader := header.Get("Authorization")
			if authHeader == "" {
				return nil, errors.New("authorization header required")
			}

			bearerToken := strings.Split(authHeader, " ")
			if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
				return nil, errors.New("invalid token format")
			}

			tokenStr := bearerToken[1]

			// Parse and validate MFA token (now returns purpose too)
			userID, email, isSetup, purpose, mfaUuid, err := jwtLib.ParseMfaToken("Bearer " + tokenStr)
			if err != nil {
				ctxSess.ErrorMessage = fmt.Sprintf("invalid MFA token: %s", err.Error())
				return nil, fmt.Errorf("invalid MFA token: %s", err.Error())
			}

			// Store purpose in context for later use
			ctxSess.Put("mfa_purpose", purpose)

			ctxSess.UserSession = models.UserSession{
				UserId: int64(userID),
				Email:  email,
			}

			// Set MFA setup status in the context map
			ctxSess.Put("mfa_setup", isSetup)
			ctxSess.Put("user_id", userID)

			if err := jwtLib.DeleteMfaAuth(ctx, mfaUuid); err != nil {
				logger.LogDebug("Invalid token:", err.Error())
			}

			logger.LogDebug("MFA middleware successful for user:", userID, "purpose:", purpose)
			return next(ctx, request)
		}
	}
}

// Helper function to wrap multiple middlewares
func WrapMiddlewares(middlewares ...endpoint.Middleware) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}

// Specific middleware combinations
func AuthWithMfaMiddleware() endpoint.Middleware {
	return WrapMiddlewares(
		MfaMiddleware(),
	)
}

func MfaOnlyMiddleware() endpoint.Middleware {
	return WrapMiddlewares(
		MfaMiddleware(),
	)
}

func HasPermission(ctxSess *ctxDM.Context, required ...string) bool {
	// If permissions are already in the session, use them
	if len(ctxSess.UserSession.Permissions) > 0 {
		for _, perm := range ctxSess.UserSession.Permissions {
			for _, req := range required {
				if perm == req {
					return true
				}
			}
		}
		return false
	}

	// Otherwise, fetch from Redis cache using user ID
	userID := ctxSess.UserSession.UserId
	if userID == 0 {
		return false
	}

	ctx := context.Background()
	_, permissions, err := jwtLib.GetCachedUserPermissions(ctx, int(userID))
	if err != nil {
		logger.LogDebug("permission.fetch.error", err.Error())
		return false
	}

	// Cache permissions in session for this request
	ctxSess.UserSession.Permissions = permissions

	for _, perm := range permissions {
		for _, req := range required {
			if perm == req {
				return true
			}
		}
	}

	return false
}
