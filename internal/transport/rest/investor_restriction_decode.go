package rest

import (
	"context"
	domain "dm_loanservice/internal/service/domain/investor_restriction"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func decodeInvestorRestrictionAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		err error
		req domain.InvestorRestriction
	)

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.WithMessage(err, "error decode request")
	}

	return req, nil
}

func decodeInvestorRestrictionReadRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := domain.InvestorRestrictionReadRequest{}
	vars := mux.Vars(r)
	req.ID = vars["id"]

	return req, nil
}
