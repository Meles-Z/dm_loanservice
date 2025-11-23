package rest

import (
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeAccountLockRuleAddHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeAccountLockRuleAddRequest, encodeDefaultResponse, opts...)
}

func MakeAccountLockRuleReadHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeAccountLockRuleReadRequest, encodeDefaultResponse, opts...)
}
