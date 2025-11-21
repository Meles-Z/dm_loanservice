package jwt

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"dm_loanservice/drivers/goconf"

	"dm_loanservice/drivers/logger"
	"dm_loanservice/drivers/models"
	redisLib "dm_loanservice/drivers/redis"
	"dm_loanservice/drivers/utils"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/twinj/uuid"
)

var (
	failedPrepareToken  = "Failed when preparing token"
	failedParseToken    = "Parse token failed"
	failedSigningMethod = "Invalid jwt signing method"
	accessKey           = goconf.Config().GetString("jwt.key_access")
	refreshKey          = goconf.Config().GetString("jwt.key_refresh")
	algo                = goconf.Config().GetString("jwt.algorithm")
	signingMethod       = jwt.GetSigningMethod(algo)
	client              = redisLib.GetConnection(context.Background())
)

var expiry time.Duration

func SetExpiry(duration int) {
	expiry = time.Duration(duration)
	return
}

type tokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
	Algo         jwt.SigningMethod
}

func NewTokenDetails(email string, id, roleID, teamID string) (*tokenDetails, error) {
	atExpires := time.Now().Add(time.Minute * expiry).Unix()
	accessUuid := uuid.NewV4().String()
	rtExpires := time.Now().Add(time.Hour * 24 * 7).Unix()
	refreshUuid := accessUuid + "++" + email

	td := &tokenDetails{}
	td.AtExpires = atExpires
	td.AccessUuid = accessUuid
	td.RtExpires = rtExpires
	td.RefreshUuid = refreshUuid

	var err error

	// Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = accessUuid
	atClaims["email"] = email
	atClaims["userId"] = id
	atClaims["exp"] = atExpires
	atClaims["rid"] = roleID
	atClaims["tid"] = teamID
	at := jwt.NewWithClaims(signingMethod, atClaims)
	td.AccessToken, err = at.SignedString([]byte(goconf.Config().GetString("jwt.key_access")))
	if err != nil {
		return nil, errors.New(failedSigningMethod + " while creating Access Token")
	}

	// Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["email"] = email
	rtClaims["userId"] = id
	rtClaims["exp"] = td.RtExpires
	rtClaims["rid"] = roleID
	rtClaims["tid"] = teamID
	rt := jwt.NewWithClaims(signingMethod, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(goconf.Config().GetString("jwt.key_refresh")))
	if err != nil {
		return nil, errors.New(failedSigningMethod + " while creating Refresh Token")
	}
	return td, nil
}
func NewMfaToken(userID int, email string, isSetup bool, purpose string) (string, string, int64, error) {
	expires := time.Now().Add(time.Minute * 10).Unix()
	mfaUuid := uuid.NewV4().String()

	claims := jwt.MapClaims{
		"mfa_uuid":  mfaUuid,
		"user_id":   userID,
		"email":     email,
		"mfa_setup": isSetup,
		"exp":       expires,
		"purpose":   purpose,
	}

	token := jwt.NewWithClaims(signingMethod, claims)
	mfaToken, err := token.SignedString([]byte(goconf.Config().GetString("mfasecret.mfasecret")))
	if err != nil {
		return "", "", 0, errors.New("failed to create MFA token")
	}

	return mfaToken, mfaUuid, expires, nil
}

func ParseAccessToken(tokens string) (*jwt.Token, error) {
	parts := strings.SplitN(tokens, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil, errors.New(failedPrepareToken)
	}

	token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
		if signingMethod != token.Method {
			return nil, errors.New(failedParseToken)
		}
		return []byte(accessKey), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func ParseRefreshToken(tokens string) (*jwt.Token, error) {
	parts := strings.SplitN(tokens, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil, errors.New(failedPrepareToken)
	}

	token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
		if signingMethod != token.Method {
			return nil, errors.New(failedParseToken)
		}
		return []byte(refreshKey), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func ParseMfaToken(tokenString string) (int, string, bool, string, string, error) {

	if !strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = "Bearer " + tokenString
	}

	parts := strings.SplitN(tokenString, " ", 2)
	if len(parts) != 2 {
		return 0, "", false, "", "", errors.New("failed to prepare MFA token")
	}

	token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
		logger.LogDebug("Token signing method:", token.Method.Alg())
		if signingMethod != token.Method {
			logger.LogDebug("Signing method mismatch. Expected:", signingMethod, "Got:", token.Method)
			return nil, errors.New("failed to parse MFA token - signing method mismatch")
		}
		return []byte(goconf.Config().GetString("mfasecret.mfasecret")), nil
	})

	if err != nil {
		logger.LogDebug("JWT parse error:", err.Error())
		return 0, "", false, "", "", err
	}

	if !token.Valid {
		logger.LogDebug("Token is invalid")
		return 0, "", false, "", "", errors.New("invalid MFA token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		logger.LogDebug("Failed to extract claims")
		return 0, "", false, "", "", errors.New("invalid MFA token claims")
	}

	// Extract custom claims
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		logger.LogDebug("user_id missing or wrong type in MFA token")
		return 0, "", false, "", "", errors.New("user_id missing in MFA token")
	}

	email, ok := claims["email"].(string)
	if !ok {
		logger.LogDebug("email missing or wrong type in MFA token")
		return 0, "", false, "", "", errors.New("email missing in MFA token")
	}

	setup, ok := claims["mfa_setup"].(bool)
	if !ok {
		logger.LogDebug("mfa_setup missing or wrong type in MFA token")
		return 0, "", false, "", "", errors.New("mfa_setup missing in MFA token")
	}

	mfaUuid, ok := claims["mfa_uuid"].(string)
	if !ok {
		logger.LogDebug("mfa_uuid missing or wrong type in MFA token")
		return 0, "", false, "", "", errors.New("mfa_uuid missing in MFA token")
	}

	purpose, ok := claims["purpose"].(string)
	if !ok {
		logger.LogDebug("purpose missing or wrong type in MFA token")
		return 0, "", false, "", "", errors.New("purpose missing in MFA token")
	}

	userID := int(userIDFloat)

	logger.LogDebug("Successfully parsed MFA token for user:", userID, "email:", email, "purpose:", purpose)

	return userID, email, setup, purpose, mfaUuid, nil
}

func NewExtractTokenMetadata(token string) (user models.UserSession, err error) {
	jwtParse, err := ParseAccessToken(token)
	if err != nil {
		logger.LogDebug("CallingHelper=", "jwt.NewExtractTokenMetadata.ParseAccessToken", "Err=", err.Error())
		err = errors.New(failedParseToken)
		return
	}
	claims, ok := jwtParse.Claims.(jwt.MapClaims)
	if ok && jwtParse.Valid {
		user.AccessUUID = claims["access_uuid"].(string)
		user.Authorized = claims["authorized"].(bool)
		user.Email = claims["email"].(string)
		user.Exp = claims["exp"].(float64)
		userID, errs := utils.Decrypt(claims["userId"].(string))
		if errs != nil {
			logger.LogDebug("CallingHelper=", "jwt.NewExtractTokenMetadata.Decrypt", "Err=", errs.Error())
			return models.UserSession{}, errors.New("failed extract metadata UserID")
		}
		user.UserId = int64(userID)
		if claims["rid"] == nil {
			logger.LogDebug("CallingHelper=", "jwt.NewExtractTokenMetadata.Decrypt", "Err=", "empty role id in token, re-login")
			return models.UserSession{}, errors.New("failed extract metadata UserID")
		}
		if claims["tid"] == nil {
			logger.LogDebug("CallingHelper=", "jwt.NewExtractTokenMetadata.Decrypt", "Err=", "empty role id in token, re-login")
			return models.UserSession{}, errors.New("failed extract metadata UserID")
		}
		roleID, errs := utils.Decrypt(claims["rid"].(string))
		if errs != nil {
			logger.LogDebug("CallingHelper=", "jwt.NewExtractTokenMetadata.Decrypt", "Err=", errs.Error())
			return models.UserSession{}, errors.New("failed extract metadata roleID")
		}
		user.RoleId = int64(roleID)
		user.TeamId = claims["tid"].(string)
	}
	return
}

func ExtractTokenMetadata(token string) (string, string, error) {
	jwtParse, err := ParseAccessToken(token)
	if err != nil {
		return "", "", errors.New(failedParseToken)
	}
	claims, ok := jwtParse.Claims.(jwt.MapClaims)
	if ok && jwtParse.Valid {
		accessUuid, errUuid := claims["access_uuid"].(string)
		if !errUuid {
			return "", "", errors.New("failed extract metadata accessUuid")
		}
		email, errEmail := claims["email"].(string)
		if !errEmail {
			return "", "", errors.New("failed extract metadata email")
		}
		return accessUuid, email, nil
	}
	return "", "", nil
}

func DeleteTokens(ctx context.Context, accessUuid, email string) error {
	// get the refresh uuid
	refreshUuid := fmt.Sprintf("%s++%s", accessUuid, email)

	// delete access token
	deletedAt, err := client.Del(ctx, accessUuid).Result()
	if err != nil {
		return errors.New("failed delete access token")
	}

	// delete refresh token
	deletedRt, err := client.Del(ctx, refreshUuid).Result()
	if err != nil {
		return errors.New("failed delete refresh token")
	}

	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong when delete access or refresh token")
	}

	return nil
}

func CreateAuth(ctx context.Context, accessUuid, refreshUuid, email string, atExpires, rtExpires int64) error {
	at := time.Unix(atExpires, 0)
	rt := time.Unix(rtExpires, 0)
	now := time.Now()

	errAccess := client.Set(ctx, accessUuid, email, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	errRefresh := client.Set(ctx, refreshUuid, email, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}

	return nil
}

func FetchAuth(ctx context.Context, accessUuid string) (string, error) {
	email, err := client.Get(ctx, accessUuid).Result()
	if err != nil {
		return "", errors.New("failed fetch auth. Invalid Authentication")
	}
	return email, nil
}

func DeleteAuth(ctx context.Context, givenUuid string) (int64, error) {
	deleted, err := client.Del(ctx, givenUuid).Result()
	if err != nil {
		return 0, errors.New("failed delete auth")
	}

	return deleted, nil
}

func CreateMfaAuth(ctx context.Context, mfaUuid, email string, expires int64) error {
	expTime := time.Unix(expires, 0)
	now := time.Now()

	err := client.Set(ctx, "mfa:"+mfaUuid, email, expTime.Sub(now)).Err()
	if err != nil {
		return errors.New("failed to store MFA token")
	}
	return nil
}

func ValidateMfaToken(ctx context.Context, mfaUuid string) (string, error) {
	email, err := client.Get(ctx, "mfa:"+mfaUuid).Result()
	if err != nil {
		return "", errors.New("MFA token not found or expired")
	}
	return email, nil
}

func DeleteMfaAuth(ctx context.Context, mfaUuid string) error {
	deleted, err := client.Del(ctx, "mfa:"+mfaUuid).Result()
	if err != nil {
		return errors.New("failed to delete MFA token")
	}
	if deleted != 1 {
		return errors.New("MFA token not found")
	}
	return nil
}

func CacheUserPermissions(ctx context.Context, userID int, roleID int, permissions []string, ttl time.Duration) error {
	data := map[string]interface{}{
		"role_id":     roleID,
		"permissions": permissions,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal permissions: %w", err)
	}

	key := fmt.Sprintf("user_permissions:%d", userID)
	err = client.Set(ctx, key, jsonData, ttl).Err()
	if err != nil {
		return fmt.Errorf("failed to cache permissions: %w", err)
	}
	return nil
}

func GetCachedUserPermissions(ctx context.Context, userID int) (int, []string, error) {
	key := fmt.Sprintf("user_permissions:%d", userID)
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, nil, nil
		}
		return 0, nil, fmt.Errorf("failed to get permissions: %w", err)
	}

	var data struct {
		RoleID      int      `json:"role_id"`
		Permissions []string `json:"permissions"`
	}

	if err := json.Unmarshal([]byte(val), &data); err != nil {
		return 0, nil, fmt.Errorf("failed to unmarshal cached permissions: %w", err)
	}

	return data.RoleID, data.Permissions, nil
}

// DeleteCachedUserPermissions removes cached permissions for a user
func DeleteCachedUserPermissions(ctx context.Context, userID int) error {
	key := fmt.Sprintf("user_permissions:%d", userID)
	err := client.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete cached permissions: %w", err)
	}
	return nil
}

// CacheDashboardMetrics caches summary metrics like arrears data.
func CacheDashboardMetrics(ctx context.Context, key string, data interface{}, ttl time.Duration) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal dashboard data: %w", err)
	}

	err = client.Set(ctx, "dashboard:"+key, jsonData, ttl).Err()
	if err != nil {
		return fmt.Errorf("failed to cache dashboard data: %w", err)
	}

	return nil
}

// GetCachedDashboardMetrics retrieves cached dashboard data.
func GetCachedDashboardMetrics[T any](ctx context.Context, key string) (*T, error) {
	val, err := client.Get(ctx, "dashboard:"+key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil // cache miss
		}
		return nil, fmt.Errorf("failed to fetch dashboard cache: %w", err)
	}

	var data T
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cached dashboard data: %w", err)
	}

	return &data, nil
}
