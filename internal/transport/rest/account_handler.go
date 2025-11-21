package rest

import (
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeAccountAddHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeAccountCreateRequest, encodeDefaultResponse, opts...)
}

func MakeAccountReadHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeAccountReadRequest, encodeDefaultResponse, opts...)
}

func MakeAccountRecentArrearsHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeAccountRecentArrearsRequest, encodeDefaultResponse, opts...)
}

func MakeAccountMortgagePerformanceHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeMortgagePerformanceRequest, encodeDefaultResponse, opts...)
}
