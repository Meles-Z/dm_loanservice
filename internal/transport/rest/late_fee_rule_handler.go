package rest

import (
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeLateFeeRuleAddHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeLateFeeRuleAddRequest, encodeDefaultResponse, opts...)
}

func MakeLateFeeRuleReadHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeLateFeeRuleReadRequest, encodeDefaultResponse, opts...)
}

func MakeLateFeeRuleUpdateHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeLateFeeRuleUpdateRequest, encodeDefaultResponse, opts...)
}
