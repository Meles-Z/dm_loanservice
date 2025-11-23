package endpoint

import (
	"context"
	ctxDM "dm_loanservice/drivers/utils/context"
	accountlockrule "dm_loanservice/internal/service/domain/account_lock_rule"
	accountlockruleSvc "dm_loanservice/internal/service/usecase/account_lock_rule"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

func makeAccountLockRuleAddEndpoint(svc accountlockruleSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(accountlockrule.AccountLockRuleAddRequest)
		if !ok {
			err := errors.New("error parse  AccountLockRule Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.AccountLockRuleAdd(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}
func makeAccountLockRuleReadEndpoint(svc accountlockruleSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(accountlockrule.AccountLockRuleReadRequest)
		if !ok {
			err := errors.New("error parse  AccountLockRule Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.AccountLockRuleRead(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}
