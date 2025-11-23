package rest

import (
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeInvestorRestrictionAddHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeInvestorRestrictionAddRequest, encodeDefaultResponse, opts...)
}

func MakeInvestorRestrictionReadHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeInvestorRestrictionReadRequest, encodeDefaultResponse, opts...)
}
