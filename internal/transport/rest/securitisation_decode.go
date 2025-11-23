package rest

import (
	"context"
	"dm_loanservice/internal/service/domain/securitisation"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func decodeEligibleAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		err error
		req securitisation.EligibleLoansQuery
	)

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.WithMessage(err, "error decode request")
	}

	return req, nil
}

func decodeEligibleAccountSummaryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeEligibleAccountSummaryReportRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		err error
		req securitisation.EligibleLoansQuery
	)

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.WithMessage(err, "error decode request")
	}
	return req, nil
}
