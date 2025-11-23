package rest

import (
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeServiceRestrictionAddHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeServiceRestrictionAddRequest, encodeDefaultResponse, opts...)
}

func MakeServiceRestrictionReadHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeServiceRestrictionReadRequest, encodeDefaultResponse, opts...)
}

func MakeServiceRestrictionReadByAccountHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeServiceRestrictionReadByAccountRequest, encodeDefaultResponse, opts...)
}
