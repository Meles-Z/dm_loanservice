package endpoint

import (
	"context"
	ctxDM "dm_loanservice/drivers/utils/context"
	dashboardSvc "dm_loanservice/internal/service/usecase/dashboard"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

func makeDashboardEndpoint(svc dashboardSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		ctxSess.Lv1("Incoming message")

		fmt.Println("We have reached here in portifolio summary")
		respOK, respErr := svc.PortfolioSummary(ctx, ctxSess)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}
