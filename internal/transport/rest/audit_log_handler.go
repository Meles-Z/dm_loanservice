package rest

import (
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeAccountAuditLogReadHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeAccountAuditLogReadRequest, encodeDefaultResponse, opts...)
}
