package endpoint

import (
	accountSvc "dm_loanservice/internal/service/usecase/account"
	duediligenceSvc "dm_loanservice/internal/service/usecase/due_diligence"
	lateFeeRuleSvc "dm_loanservice/internal/service/usecase/late_fee_rule"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {

	// account
	AccountAdd          endpoint.Endpoint
	AccountRead         endpoint.Endpoint
	RecentArrears       endpoint.Endpoint
	MortgagePerformance endpoint.Endpoint

	// due diligence
	DueDiligenceAdd       endpoint.Endpoint
	DueDiligenceRead      endpoint.Endpoint
	DueDiligenceUpdate    endpoint.Endpoint
	DueDiligenceByAccount endpoint.Endpoint

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
	duediligenceSvc duediligenceSvc.Service,

) Endpoints {
	return Endpoints{

		// Account
		AccountAdd:          makeAccountAddEndpoint(accountSvc),
		AccountRead:         makeAccountReadEndpoint(accountSvc),
		RecentArrears:       makeAccountRecentArrearsEndpoint(accountSvc),
		MortgagePerformance: makeAccountMortgagePerformanceEndpoint(accountSvc),

		// Due Diligence
		DueDiligenceAdd:       makeDueDiligenceAddEndpoint(duediligenceSvc),
		DueDiligenceRead:      makeDueDiligenceReadEndpoint(duediligenceSvc),
		DueDiligenceUpdate:    makeDueDiligenceUpdateEndpoint(duediligenceSvc),
		DueDiligenceByAccount: makeDueDiligenceByAccountEndpoint(duediligenceSvc),

		LateFeeRuleAdd:    makeLateFeeRuleAddEndpoint(lateFeeRuleSvc),
		LateFeeRuleRead:   makeLateFeeRuleReadEndpoint(lateFeeRuleSvc),
		LateFeeRuleUpdate: makeLateFeeRuleUpdateEndpoint(lateFeeRuleSvc),
	}
}
