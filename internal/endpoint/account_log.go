package endpoint

import (
	"context"
	ctxDM "dm_loanservice/drivers/utils/context"
	accountauditlog "dm_loanservice/internal/service/domain/account_audit_log"
	accountAuditLogSvc "dm_loanservice/internal/service/usecase/account_audit_log"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

func makeAccountAuditLogReadEndpoint(svc accountAuditLogSvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		data := ctx.Value(ctxDM.AppSession)
		ctxSess := data.(*ctxDM.Context)
		r, ok := req.(accountauditlog.AccountAuditLogReadRequest)
		if !ok {
			err := errors.New("error parse  AccountAuditLog Request")
			ctxSess.Lv4(err)
			return nil, err
		}
		ctxSess.Request = r
		ctxSess.Lv1("Incoming message")

		respOK, respErr := svc.AccountAuditLogRead(ctx, ctxSess, r)
		if respErr != nil {
			ctxSess.Lv4(respErr)
			return respErr, nil
		}
		ctxSess.Lv4(respOK)
		return respOK, nil
	}
}
