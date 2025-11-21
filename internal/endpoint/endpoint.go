package endpoint

import (
	lateFeeRuleSvc "dm_loanservice/internal/service/usecase/late_fee_rule"
	accountSvc "dm_loanservice/internal/service/usecase/account"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {

	// account
	AccountAdd          endpoint.Endpoint
	AccountRead         endpoint.Endpoint
	RecentArrears       endpoint.Endpoint
	MortgagePerformance endpoint.Endpoint

	// product
	ProductAdd                  endpoint.Endpoint
	ProductSearch               endpoint.Endpoint
	ProductRead                 endpoint.Endpoint
	ProductUpdate               endpoint.Endpoint
	ProductDelete               endpoint.Endpoint
	ProductAnalysisCount        endpoint.Endpoint
	ProductOverpaymentLimit     endpoint.Endpoint
	ProductOverpaymentCalculate endpoint.Endpoint

	// lateFeeRule
	LateFeeRuleAdd    endpoint.Endpoint
	LateFeeRuleRead   endpoint.Endpoint
	LateFeeRuleUpdate endpoint.Endpoint
}

func NewEndpoints(
	lateFeeRuleSvc lateFeeRuleSvc.Service,
	accountSvc accountSvc.Service,

) Endpoints {
	return Endpoints{

		// Account
		AccountAdd:          makeAccountAddEndpoint(accountSvc),
		AccountRead:         makeAccountReadEndpoint(accountSvc),
		RecentArrears:       makeAccountRecentArrearsEndpoint(accountSvc),
		MortgagePerformance: makeAccountMortgagePerformanceEndpoint(accountSvc),

		LateFeeRuleAdd:    makeLateFeeRuleAddEndpoint(lateFeeRuleSvc),
		LateFeeRuleRead:   makeLateFeeRuleReadEndpoint(lateFeeRuleSvc),
		LateFeeRuleUpdate: makeLateFeeRuleUpdateEndpoint(lateFeeRuleSvc),
	}
}
