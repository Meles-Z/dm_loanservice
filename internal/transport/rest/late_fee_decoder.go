package rest

import (
	"context"
	domain "dm_loanservice/internal/service/domain/late_fee_rule"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func decodeLateFeeRuleAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		err error
		req domain.LateFeeRuleAddRequest
	)

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.WithMessage(err, "error decode request")
	}

	return req, nil
}

func decodeLateFeeRuleReadRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := domain.LateFeeRuleReadRequest{}
	vars := mux.Vars(r)
	req.Id = vars["id"]

	return req, nil
}

func decodeLateFeeRuleUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		err error
		req domain.LateFeeRuleUpdateRequest
	)

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.WithMessage(err, "error decode request")
	}

	return req, nil
}
