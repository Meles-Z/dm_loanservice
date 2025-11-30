package rest

import (
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeTaskAddHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeTaskAddRequest, encodeDefaultResponse, opts...)
}

func MakeRecentTaskHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeRecentTaskRequest, encodeDefaultResponse, opts...)
}

func MakeTaskSummaryHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeTaskSummaryRequest, encodeDefaultResponse, opts...)
}
