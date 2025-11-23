package rest

import (
	"context"
	domain "dm_loanservice/internal/service/domain/account"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func decodeAccountCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		err error
		req domain.AccountAddRequest
	)
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.WithMessage(err, "error decode request")
	}

	return req, nil
}

func decodeAccountReadRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := domain.AccountReadRequest{}
	vars := mux.Vars(r)
	req.ID = vars["id"]

	return req, nil
}

func decodeAccountUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		err error
		req domain.AccountUpdateRequest
	)

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.WithMessage(err, "error decode request")
	}

	vars := mux.Vars(r)
	req.ID = vars["id"]
	return req, nil
}

func decodeAccountRecentArrearsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// var (
	// 	err error
	// 	req domain.AccountRequest
	// )
	// err = json.NewDecoder(r.Body).Decode(&req)
	// if err != nil {
	// 	return nil, errors.WithMessage(err, "error decode request")
	// }

	return nil, nil
}

func decodeMortgagePerformanceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
