package rest

import (
	"context"
	"dm_loanservice/internal/service/domain/securitisation"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

func decodeSecuritisationPoolAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		err error
		req securitisation.SecuritisationPoolAddRequest
	)

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.WithMessage(err, "error decode request")
	}
	return req, nil
}

func decodeSecuritisationPoolReadRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := securitisation.SecuritisationPoolReadRequest{}
	vars := mux.Vars(r)
	req.ID = vars["id"]
	return req, nil
}

func decodeSecuritisationPoolUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		err error
		req securitisation.SecuritisationPoolUpdateRequest
	)

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.WithMessage(err, "error decode request")
	}
	return req, nil
}

func decodeSecuritisationDeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := securitisation.SecuritisationDeleteRequest{}
	vars := mux.Vars(r)
	req.ID = vars["id"]
	return req, nil
}

func decodeSecuritisationPoolAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeSecuritisationDashboardRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeSecuritisationDashboardExportRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		err error
		req securitisation.DashboardExportRequest
	)

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.WithMessage(err, "error decode request")
	}

	return req, nil
}

func decodeSecuritisationPoolReportRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		err error
		req securitisation.SecuritisationPoolReportRequest
	)

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.WithMessage(err, "error decode request")
	}
	vars := mux.Vars(r)
	req.ID = vars["id"]
	return req, nil
}
