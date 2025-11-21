package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	ctxDM "dm_loanservice/drivers/utils/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const (
	CtxSession     = "Context_session"
	XCorrelationID = "X-Correlation-ID"
	Header         = "header"
	Authorization  = "Authorization"
)

type RpcConnection struct {
	Conn    *grpc.ClientConn
	options Options
}

type Options struct {
	PathCredential string
	ServerName     string
	Address        string
	Timeout        time.Duration
}

func NewGrpcConnection(options Options) *RpcConnection {
	var conn *grpc.ClientConn
	var err error

	if len(options.PathCredential) == 0 {
		conn, err = grpc.Dial(options.Address, grpc.WithInsecure(), withClientUnaryInterceptor())
	} else {
		creds, errs := credentials.NewClientTLSFromFile(options.PathCredential, options.ServerName)
		if errs != nil {
			log.Fatalf("could not process the credentials: %v", errs)
		}

		conn, err = grpc.Dial(options.Address, grpc.WithTransportCredentials(creds), withClientUnaryInterceptor())
	}
	if err != nil {
		panic(err)
	}

	return &RpcConnection{
		Conn:    conn,
		options: options,
	}
}

func (rpc *RpcConnection) CreateContext(parentCtx context.Context, ctxSess *ctxDM.Context) (ctx context.Context) {
	ctx, _ = context.WithTimeout(parentCtx, rpc.options.Timeout*time.Second)
	ctx = context.WithValue(ctx, CtxSession, ctxSess)
	header, _ := json.Marshal(ctxSess.Header)

	ctx = metadata.AppendToOutgoingContext(ctx,
		XCorrelationID, ctxSess.XCorrelationID,
		Header, string(header),
	)

	return
}

func clientInterceptor(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
	ctxSess := ctx.Value(CtxSession).(*ctxDM.Context)
	auth := ""
	if head, ok := ctxSess.Header.(http.Header); ok { // get auth from rest call
		auth = head.Get(Authorization)
	}
	if ctxSess.Method == "GRPC" { // get auth from grpc call
		auth = fmt.Sprintf("Bearer %s", ctxSess.GrpcAuthToken)
	}
	ctxWithMetadata := metadata.NewOutgoingContext(ctx,
		metadata.Pairs(
			XCorrelationID, ctxSess.XCorrelationID,
			Authorization, auth,
		),
	)

	err = invoker(ctxWithMetadata, method, req, reply, cc, opts...)

	return
}

func withClientUnaryInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(clientInterceptor)
}

