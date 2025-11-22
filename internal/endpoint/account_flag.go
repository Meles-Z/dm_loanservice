package endpoint

import (
	"context"
	ctxDM "dm_loanservice/drivers/utils/context"
	accountflag "dm_loanservice/internal/service/domain/account_flag"
	accountflagSvc "dm_loanservice/internal/service/usecase/account_flag"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

func makeAccountFlagAddEndpoint(svc accountflagSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(accountflag.AccountFlagAddRequest)
		if !ok {
			err := errors.New("error parse  AccountFlag Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.AccountFlagAdd(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}
func makeAccountFlagReadEndpoint(svc accountflagSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(accountflag.AccountFlagReadRequest)
		if !ok {
			err := errors.New("error parse  AccountFlag Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.AccountFlagRead(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}
