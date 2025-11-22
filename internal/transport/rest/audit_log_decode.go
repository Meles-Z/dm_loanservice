package rest

import (
	"context"
	"net/http"

	domain "dm_loanservice/internal/service/domain/account_audit_log"

	"github.com/gorilla/mux"
)

func decodeAccountAuditLogReadRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := domain.AccountAuditLogReadRequest{}
	vars := mux.Vars(r)
	req.ID = vars["id"]

	return req, nil
}
