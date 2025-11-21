package endpoint

import (
	"context"
	ctxDM "dm_loanservice/drivers/utils/context"
	latefeerule "dm_loanservice/internal/service/domain/late_fee_rule"
	lateFeeRuleSvc "dm_loanservice/internal/service/usecase/late_fee_rule"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

func makeLateFeeRuleAddEndpoint(svc lateFeeRuleSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(latefeerule.LateFeeRuleAddRequest)
		if !ok {
			err := errors.New("error parse PaymentGetByIDRequest")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.LateFeeRuleAdd(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}

		ctxSess.Response = respOK
		ctxSess.Lv4()

		return respOK, nil
	}
}

func makeLateFeeRuleReadEndpoint(svc lateFeeRuleSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(latefeerule.LateFeeRuleReadRequest)
		if !ok {
			err := errors.New("error parse PaymentGetByIDRequest")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.LateFeeRuleRead(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}

		ctxSess.Response = respOK
		ctxSess.Lv4()

		return respOK, nil
	}
}

func makeLateFeeRuleUpdateEndpoint(svc lateFeeRuleSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(latefeerule.LateFeeRuleUpdateRequest)
		if !ok {
			err := errors.New("error parse PaymentGetByIDRequest")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.LateFeeRuleUpdate(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}

		ctxSess.Response = respOK
		ctxSess.Lv4()

		return respOK, nil
	}
}
