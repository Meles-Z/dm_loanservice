package rest

import (
	"context"
	domain "dm_loanservice/internal/service/domain/service_restriction"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func decodeServiceRestrictionAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		err error
		req domain.ServiceRestrictionAddRequest
	)

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.WithMessage(err, "error decode request")
	}

	return req, nil
}

func decodeServiceRestrictionReadRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := domain.ServiceRestrictionReadRequest{}
	vars := mux.Vars(r)
	req.ID = vars["id"]

	return req, nil
}


func decodeServiceRestrictionReadByAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := domain.ServiceRestrictionReadByAccountRequest{}
	vars := mux.Vars(r)
	req.ID = vars["id"]

	return req, nil
}

