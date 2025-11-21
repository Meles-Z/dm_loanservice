package rest

import (
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeDueDiligenceAddHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeDueDiligenceAddRequest, encodeDefaultResponse, opts...)
}
func MakeDueDiligenceReadHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeDueDiligenceReadRequest, encodeDefaultResponse, opts...)
}

func MakeDueDiligenceUpdateHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeDueDiligenceUpdateRequest, encodeDefaultResponse, opts...)
}
func MakeDueDiligenceByAccountHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeDueDiligenceByAccountRequest, encodeDefaultResponse, opts...)
}
