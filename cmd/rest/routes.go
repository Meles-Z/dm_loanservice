package rest

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"dm_loanservice/drivers/middleware"
	ie "dm_loanservice/internal/endpoint"
	ithttp "dm_loanservice/internal/transport/rest"
)

// aliveHandler is a simple health check endpoint
func aliveHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok","message":"service is alive"}`))
}

func initRoutes(ctx context.Context, e ie.Endpoints) *mux.Router {
	r := mux.NewRouter()

	// account
	accountRoutes := r.PathPrefix("/api/v1/account").Subrouter()
	accountRoutes.Handle("", ithttp.MakeAccountAddHandler((e.AccountAdd))).Methods(http.MethodPost)
	accountRoutes.Handle("/{id:[0-9a-fA-F-]+}", ithttp.MakeAccountReadHandler((e.AccountRead))).Methods(http.MethodGet)
	accountRoutes.Handle("/arrears/recent", ithttp.MakeAccountRecentArrearsHandler(e.RecentArrears)).Methods(http.MethodGet)
	accountRoutes.Handle("/analytics/mortgage-performance", ithttp.MakeAccountMortgagePerformanceHandler(e.MortgagePerformance)).Methods(http.MethodGet)
	accountRoutes.Handle("/{id:[0-9a-fA-F-]+}/due-diligence", ithttp.MakeDueDiligenceByAccountHandler((e.DueDiligenceByAccount))).Methods(http.MethodGet)
	accountRoutes.Handle("/{id:[0-9a-fA-F-]+}/due-diligence", ithttp.MakeDueDiligenceUpdateHandler((e.DueDiligenceUpdate))).Methods(http.MethodPut)
	accountRoutes.Handle("/{id:[0-9a-fA-F-]+}/flag", ithttp.MakeAccountFlagAddHandler((e.AccountFlagAdd))).Methods(http.MethodPost)
	accountRoutes.Handle("/{id:[0-9a-fA-F-]+}/audit-log", ithttp.MakeAccountAuditLogReadHandler((e.AccountAuditLogRead))).Methods(http.MethodGet)

	// due diligence
	duediligenceRoutes := r.PathPrefix("/api/v1/due-diligence").Subrouter()
	duediligenceRoutes.Handle("/{id:[0-9a-fA-F-]+}", ithttp.MakeDueDiligenceReadHandler((e.DueDiligenceRead))).Methods(http.MethodGet)

	// account flag
	accountFlagRoutes := r.PathPrefix("/api/v1/account-flag").Subrouter()
	accountFlagRoutes.Handle("", ithttp.MakeAccountFlagAddHandler((e.AccountFlagAdd))).Methods(http.MethodPost)
	accountFlagRoutes.Handle("/{id:[0-9a-fA-F-]+}", ithttp.MakeAccountFlagReadHandler((e.AccountFlagRead))).Methods(http.MethodGet)
	// late fee rule
	lateFeeRuleRouter := r.PathPrefix("/api/v1/late-fee-rule").Subrouter()
	lateFeeRuleRouter.Handle("", ithttp.MakeLateFeeRuleAddHandler((e.LateFeeRuleAdd))).Methods(http.MethodPost)
	lateFeeRuleRouter.Handle("/{id:[0-9a-fA-F-]+}", ithttp.MakeLateFeeRuleReadHandler(middleware.AuthMiddleware()(e.LateFeeRuleRead))).Methods(http.MethodGet)
	lateFeeRuleRouter.Handle("/{id:[0-9a-fA-F-]+}", ithttp.MakeLateFeeRuleUpdateHandler(middleware.AuthMiddleware()(e.LateFeeRuleUpdate))).Methods(http.MethodPut)

	return r
}
