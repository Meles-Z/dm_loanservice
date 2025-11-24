package endpoint

import (
	"context"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/internal/service/domain/securitisation"
	securitisationSvc "dm_loanservice/internal/service/usecase/securitisation"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

func makeAccountEligibleEndpoint(svc securitisationSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(securitisation.EligibleLoansQuery)
		if !ok {
			err := errors.New("error parse  Account Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.EligibleAccount(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}

func MakeEligibleAccountSummaryHandler(svc securitisationSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		// r, ok := req.(securitisation.EligibleLoansQuery)
		// if !ok {
		// 	err := errors.New("error parse  EligibleLoansQuery")
		// 	ctxSess.Lv4(err)
		// 	return nil, err
		// }
		// ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.EligibleSummary(ctx, ctxSess)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}

func MakeEligibleAccountSummaryReportHandler(svc securitisationSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(securitisation.EligibleLoansQuery)
		if !ok {
			err := errors.New("error parse  EligibleLoansQuery")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.EligibleAccountSummaryReport(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}

func MakeSecuritisationPoolAddHandler(svc securitisationSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(securitisation.SecuritisationPoolAddRequest)
		if !ok {
			err := errors.New("error parse  SecuritisationPool Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.SecuritisationPoolAdd(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}
func MakeSecuritisationPoolReadHandler(svc securitisationSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(securitisation.SecuritisationPoolReadRequest)
		if !ok {
			err := errors.New("error parse  SecuritisationPool Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.SecuritisationPoolRead(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}

func MakeSecuritisationPoolUpdateHandler(svc securitisationSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(securitisation.SecuritisationPoolUpdateRequest)
		if !ok {
			err := errors.New("error parse  SecuritisationPool Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.SecuritisationPoolUpdate(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}

func MakeSecuritisationPoolAllHandler(svc securitisationSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		// r, ok := req.(securitisation.SecuritisationPoolRequest)
		// if !ok {
		// 	err := errors.New("Error parse  SecuritisationPool Request")
		// 	ctxSess.Lv4(err)
		// 	return nil, err
		// }
		// ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.SecuritisationPoolAll(ctx, ctxSess)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}
func MakeSecuritisationDeleteHandler(svc securitisationSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(securitisation.SecuritisationDeleteRequest)
		if !ok {
			err := errors.New("error parse  SecuritisationPool Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.SecuritisationDelete(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}

func MakeSecuritisationDashboardHandler(svc securitisationSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		// r, ok := req.(securitisation.EligibleLoansQuery)
		// if !ok {
		// 	err := errors.New("error parse  EligibleLoansQuery")
		// 	ctxSess.Lv4(err)
		// 	return nil, err
		// }
		// ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.SecuritisationDashboard(ctx, ctxSess)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}

func MakeSecuritisationDashboardExportHandler(svc securitisationSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(securitisation.DashboardExportRequest)
		if !ok {
			err := errors.New("error parse  DashboardExportRequest")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.SecuritisationDashboardExport(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}

func MakeSecuritisationPoolReportHandler(svc securitisationSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(securitisation.SecuritisationPoolReportRequest)
		if !ok {
			err := errors.New("error parse  SecuritisationPool Report Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.SecuritisationPoolReport(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}
