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

func MakeSecuritisationPoolAddHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeSecuritisationPoolAddRequest, encodeDefaultResponse, opts...)
}

func MakeSecuritisationPoolReadHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeSecuritisationPoolReadRequest, encodeDefaultResponse, opts...)
}

func MakeSecuritisationPoolUpdateHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeSecuritisationPoolUpdateRequest, encodeDefaultResponse, opts...)
}

func MakeSecuritisationPoolAllHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeSecuritisationPoolAllRequest, encodeDefaultResponse, opts...)
}

func MakeSecuritisationDeleteHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeSecuritisationDeleteRequest, encodeDefaultResponse, opts...)
}

func MakeSecuritisationDashboardHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeSecuritisationDashboardRequest, encodeDefaultResponse, opts...)
}

func MakeSecuritisationDashboardExportHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeSecuritisationDashboardExportRequest, encodeDefaultResponse, opts...)
}

func MakeSecuritisationPoolReportHandler(e endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(e, decodeSecuritisationPoolReportRequest, encodeDefaultResponse, opts...)
}