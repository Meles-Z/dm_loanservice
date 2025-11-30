package endpoint

import (
	"context"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/internal/service/domain/tasks"
	tasksSvc "dm_loanservice/internal/service/usecase/tasks"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

func makeTaskAddEndpoint(svc tasksSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(tasks.TaskAddRequest)
		if !ok {
			err := errors.New("error parse  customer Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.TaskAdd(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}


func makeRecentTaskEndpoint(svc tasksSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.RecentTasks(ctx, ctxSess)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}


func makeTaskSummaryEndpoint(svc tasksSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.TaskSummary(ctx, ctxSess)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}
