package endpoint

import (
	"context"
	ctxDM "dm_loanservice/drivers/utils/context"
	investorRestriction "dm_loanservice/internal/service/domain/investor_restriction"
	investorRestrictionSvc "dm_loanservice/internal/service/usecase/investor_restriction"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

func makeInvestorRestrictionAddEndpoint(svc investorRestrictionSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(investorRestriction.InvestorRestrictionAddRequest)
		if !ok {
			err := errors.New("error parse  InvestorRestriction Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.InvestorRestrictionAdd(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}
func makeInvestorRestrictionReadEndpoint(svc investorRestrictionSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(investorRestriction.InvestorRestrictionReadRequest)
		if !ok {
			err := errors.New("error parse  InvestorRestriction Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.InvestorRestrictionRead(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}
