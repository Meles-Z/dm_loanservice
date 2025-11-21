package endpoint

import (
	"context"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/internal/service/domain/account"
	accountAvc "dm_loanservice/internal/service/usecase/account"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

func makeAccountAddEndpoint(svc accountAvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(account.AccountAddRequest)
		if !ok {
			err := errors.New("Error parse  Account Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.AccountAdd(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}

func makeAccountReadEndpoint(svc accountAvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(account.AccountReadRequest)
		if !ok {
			err := errors.New("Error parse  Account Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.AccountId(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}

func makeAccountRecentArrearsEndpoint(svc accountAvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		// r, ok := req.(account.AccountRequest)
		// if !ok {
		// 	err := errors.New("Error parse  Account Request")
		// 	ctxSess.Lv4(err)
		// 	return nil, err
		// }
		// ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.AccountRecentArrears(ctx, ctxSess)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}

func makeAccountMortgagePerformanceEndpoint(svc accountAvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.MortgagePerformance(ctx, ctxSess)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}
