package endpoint

import (
	"context"
	ctxDM "dm_loanservice/drivers/utils/context"
	serviceRestriction "dm_loanservice/internal/service/domain/service_restriction"
	serviceRestrictionSvc "dm_loanservice/internal/service/usecase/service_restriction"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

func makeServiceRestrictionAddEndpoint(svc serviceRestrictionSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(serviceRestriction.ServiceRestrictionAddRequest)
		if !ok {
			err := errors.New("error parse  ServiceRestriction Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.ServiceRestrictionAdd(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}
func makeServiceRestrictionReadEndpoint(svc serviceRestrictionSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(serviceRestriction.ServiceRestrictionReadRequest)
		if !ok {
			err := errors.New("error parse  ServiceRestriction Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.ServiceRestrictionRead(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}

func makeServiceRestrictionReadByAccountEndpoint(svc serviceRestrictionSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(serviceRestriction.ServiceRestrictionReadByAccountRequest)
		if !ok {
			err := errors.New("error parse  ServiceRestriction Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.ServiceRestrictionReadByAccount(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}
