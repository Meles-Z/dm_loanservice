package rest

import (
	"context"
	"encoding/json"
	"net/http"

	domain "dm_loanservice/internal/service/domain/account_flag"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func decodeAccountFlagAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		err error
		req domain.AccountFlagAddRequest
	)
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.WithMessage(err, "error decode request")
	}
	vars := mux.Vars(r)
	req.AccountID = vars["id"]

	return req, nil
}

func decodeAccountFlagReadRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := domain.AccountFlagReadRequest{}
	vars := mux.Vars(r)
	req.ID = vars["id"]

	return req, nil
}
