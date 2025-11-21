package rest

import (
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeAccountFlagAddHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeAccountFlagAddRequest, encodeDefaultResponse, opts...)
}
func MakeAccountFlagReadHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeAccountFlagReadRequest, encodeDefaultResponse, opts...)
}
