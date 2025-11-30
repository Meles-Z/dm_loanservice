package grpc

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"strings"

	"dm_loanservice/drivers/goconf"
	jwtLib "dm_loanservice/drivers/jwt"
	"dm_loanservice/drivers/logger"
	Logger "dm_loanservice/drivers/logger/zap"
	"dm_loanservice/drivers/utils"
	ctxDM "dm_loanservice/drivers/utils/context"

	pbAccount "github.com/brianjobling/dm_proto/generated/accountservice/accountpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const SessionID = "Session_Id"

func RunServer(
	ctx context.Context,
	pbAccountServer pbAccount.AccountServiceServer,
	port string,
) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption

	// optional interceptor
	if goconf.Config().GetString("grpc.unary_interceptor") == "enabled" {
		opts = append(opts, grpc.UnaryInterceptor(unaryServerInterceptor))
	}
	// TLS
	if goconf.Config().GetBool("grpc.tls_enabled") {
		serverCert := goconf.Config().GetString("grpc.server_cert")
		serverKey := goconf.Config().GetString("grpc.server_key")

		pair, err := tls.X509KeyPair([]byte(serverCert), []byte(serverKey))
		if err != nil {
			panic(err)
		}

		creds := credentials.NewTLS(&tls.Config{
			Certificates: []tls.Certificate{pair},
		})

		opts = append(opts, grpc.Creds(creds))
	}

	server := grpc.NewServer(opts...)
	pbAccount.RegisterAccountServiceServer(server, pbAccountServer)

	logger.LogInfo("status-gRPC", "listening", "port", port)
	return server.Serve(listen)
}

func unaryServerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Println("--> unary interceptor: ", info.FullMethod)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "UnaryEcho: failed to get metadata")
	}

	var reqId string
	if len(md.Get("X-Correlation-ID")) > 0 {
		reqId = md.Get("X-Correlation-ID")[0]
	}
	if len(reqId) == 0 {
		reqId = utils.GenerateThreadId()
	}
	ctxSess := ctxDM.New(Logger.GetLogger()).
		SetXCorrelationID(reqId).
		SetAppName("DMortgages_operations").
		SetAppVersion("0.0").
		SetPort(50051).
		SetRequest(req).
		SetURL(info.FullMethod).
		SetMethod("GRPC")

	// Check authorization
	if info.FullMethod != "/authpb.AuthService/Login" && info.FullMethod != "/authpb.AuthService/Logout" &&
		info.FullMethod != "/authpb.AuthService/Refresh" && info.FullMethod != "/authpb.AuthService/ForgotPassword" &&
		info.FullMethod != "/authpb.AuthService/CodePassword" && info.FullMethod != "/authpb.AuthService/NewPassword" &&
		info.FullMethod != "/authpb.AuthService/GenerateOtp" && info.FullMethod != "/authpb.AuthService/ValidateOtp" &&
		info.FullMethod != "/authpb.AuthService/GenerateOtpForgotPassword" && info.FullMethod != "/authpb.AuthService/ValidateOtpForgotPassword" &&
		info.FullMethod != "/userpb.UserService/UserRead" && info.FullMethod != "/userpb.UserService/UserRegister" &&
		info.FullMethod != "/userpb.UserService/CheckEmail" {

		authMD, ok := md["authorization"]
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
		}

		if len(md.Get("authorization")) > 0 {
			authHeader := authMD[0]
			bearerToken := strings.Split(authHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwtLib.ParseAccessToken(authHeader)
				if err != nil {
					return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("UnaryEcho: %s", err.Error()))
				}

				if token.Valid {
					userSession, errExtract := jwtLib.NewExtractTokenMetadata(authHeader)
					if errExtract != nil {
						return nil, status.Errorf(codes.InvalidArgument, "UnaryEcho: failed to parse access token")
					}
					ctxSess.UserSession = userSession
					ctxSess.SetGrpcAuthToken(token.Raw)

				} else {
					return nil, status.Errorf(codes.Unauthenticated, "UnaryEcho: invalid token")
				}
			} else {
				return nil, status.Errorf(codes.InvalidArgument, "UnaryEcho: invalid token")
			}
		}
	}

	return handler(context.WithValue(ctx, ctxDM.AppSession, ctxSess), req)
}
