package rest

import (
	"context"
	"dm_loanservice/internal/service/domain/securitisation"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func decodeEligibleAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := securitisation.EligibleLoansQuery{}

	// Parse int fields safely
	req.Page, _ = strconv.Atoi(r.URL.Query().Get("page"))
	req.PageSize, _ = strconv.Atoi(r.URL.Query().Get("page_size"))
	req.ArrearsDaysMax, _ = strconv.Atoi(r.URL.Query().Get("arrears_days_max"))

	// Parse float fields
	req.LTVMin, _ = strconv.ParseFloat(r.URL.Query().Get("ltv_min"), 64)
	req.LTVMax, _ = strconv.ParseFloat(r.URL.Query().Get("ltv_max"), 64)

	// Strings (safe)
	req.MortgageType = r.URL.Query().Get("mortgage_type")
	req.Region = r.URL.Query().Get("region")
	req.OriginationFrom = r.URL.Query().Get("origination_from")
	req.OriginationTo = r.URL.Query().Get("origination_to")
	req.PropertyType = r.URL.Query().Get("property_type")
	req.SortBy = r.URL.Query().Get("sort_by")
	req.SortDirection = r.URL.Query().Get("sort_direction")

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
	vars := mux.Vars(r)
	req.ID = vars["id"]

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

	req := securitisation.DashboardExportRequest{}

	req.Format = r.URL.Query().Get("format")

	return req, nil
}

func decodeSecuritisationPoolReportRequest(_ context.Context, r *http.Request) (interface{}, error) {

	req := securitisation.SecuritisationPoolReportRequest{}

	vars := mux.Vars(r)
	req.ID = vars["id"]
	return req, nil
}
