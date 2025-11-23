package rest

import (
	"context"
	"encoding/json"
	"net/http"

	domain "dm_loanservice/internal/service/domain/account_lock_rule"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func decodeAccountLockRuleAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		err error
		req domain.AccountLockRuleAddRequest
	)

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.WithMessage(err, "error decode request")
	}

	return req, nil
}

func decodeAccountLockRuleReadRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := domain.AccountLockRuleReadRequest{}
	vars := mux.Vars(r)
	req.ID = vars["id"]

	return req, nil
}
