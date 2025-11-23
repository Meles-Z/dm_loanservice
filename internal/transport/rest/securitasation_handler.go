package rest

import (
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeEligibleAccountHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeEligibleAccountRequest, encodeDefaultResponse, opts...)
}

func MakeEligibleAccountSummaryHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeEligibleAccountSummaryRequest, encodeDefaultResponse, opts...)
}

func MakeEligibleAccountSummaryReportHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeEligibleAccountSummaryReportRequest, encodeDefaultResponse, opts...)
}
