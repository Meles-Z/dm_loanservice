package rest

import (
	"context"
	domain "dm_loanservice/internal/service/domain/due_diligence"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func decodeDueDiligenceAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		err error
		req domain.DueDiligenceAddRequest
	)
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.WithMessage(err, "error decode request")
	}

	return req, nil
}

func decodeDueDiligenceReadRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := domain.DueDiligenceReadRequest{}
	vars := mux.Vars(r)
	req.ID = vars["id"]

	return req, nil
}

func decodeDueDiligenceUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		err error
		req domain.DueDiligenceUpdateRequest
	)
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.WithMessage(err, "error decode request")
	}
	vars := mux.Vars(r)
	req.ID = vars["id"]

	return req, nil
}

func decodeDueDiligenceByAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := domain.DueDiligenceByAccountRequest{}
	vars := mux.Vars(r)
	req.AccountID = vars["id"]
	return req, nil
}
