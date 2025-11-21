package endpoint

import (
	"context"
	ctxDM "dm_loanservice/drivers/utils/context"
	duediligence "dm_loanservice/internal/service/domain/due_diligence"
	duediligenceSvc "dm_loanservice/internal/service/usecase/due_diligence"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

func makeDueDiligenceAddEndpoint(svc duediligenceSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(duediligence.DueDiligenceAddRequest)
		if !ok {
			err := errors.New("error parse  DueDiligence Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.DueDiligenceAdd(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}

func makeDueDiligenceReadEndpoint(svc duediligenceSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(duediligence.DueDiligenceReadRequest)
		if !ok {
			err := errors.New("error parse  DueDiligence Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.DueDiligenceRead(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}

func makeDueDiligenceUpdateEndpoint(svc duediligenceSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(duediligence.DueDiligenceUpdateRequest)
		if !ok {
			err := errors.New("error parse  DueDiligence Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.DueDiligenceUpdate(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}

func makeDueDiligenceByAccountEndpoint(svc duediligenceSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(duediligence.DueDiligenceByAccountRequest)
		if !ok {
			err := errors.New("error parse  DueDiligence Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.DueDiligenceByAccount(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}
