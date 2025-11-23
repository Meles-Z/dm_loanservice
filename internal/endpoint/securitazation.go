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
