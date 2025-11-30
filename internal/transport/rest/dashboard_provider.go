package rest

import (
	"fmt"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeDashboardHandler(e endpoint.Endpoint) *httptransport.Server {
	fmt.Println("passed")
	return httptransport.NewServer(e, decodeDashboardRequest, encodeDefaultResponse, opts...)
}
