package endpoint

import (
	accountSvc "dm_loanservice/internal/service/usecase/account"
	accountAuditLogSvc "dm_loanservice/internal/service/usecase/account_audit_log"
	accountflagSvc "dm_loanservice/internal/service/usecase/account_flag"
	accountLockRuleSvc "dm_loanservice/internal/service/usecase/account_lock_rule"
	dashboardSvc "dm_loanservice/internal/service/usecase/dashboard"
	duediligenceSvc "dm_loanservice/internal/service/usecase/due_diligence"
	investorRestrictionSvc "dm_loanservice/internal/service/usecase/investor_restriction"
	lateFeeRuleSvc "dm_loanservice/internal/service/usecase/late_fee_rule"
	securitisationSvc "dm_loanservice/internal/service/usecase/securitisation"
	serviceRestrictionSvc "dm_loanservice/internal/service/usecase/service_restriction"
	tasksSvc "dm_loanservice/internal/service/usecase/tasks"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {

	// account
	AccountAdd          endpoint.Endpoint
	AccountRead         endpoint.Endpoint
	AccountUpdate       endpoint.Endpoint
	RecentArrears       endpoint.Endpoint
	MortgagePerformance endpoint.Endpoint

	// due diligence
	DueDiligenceAdd       endpoint.Endpoint
	DueDiligenceRead      endpoint.Endpoint
	DueDiligenceUpdate    endpoint.Endpoint
	DueDiligenceByAccount endpoint.Endpoint

	// account flag
	AccountFlagAdd  endpoint.Endpoint
	AccountFlagRead endpoint.Endpoint

	// account audit log
	AccountAuditLogRead endpoint.Endpoint

	// dashboard
	PortfolioSummary endpoint.Endpoint

	// account lock rule
	AccountLockRuleAdd  endpoint.Endpoint
	AccountLockRuleRead endpoint.Endpoint

	// service restriction
	ServiceRestrictionAdd           endpoint.Endpoint
	ServiceRestrictionRead          endpoint.Endpoint
	ServiceRestrictionReadByAccount endpoint.Endpoint

	// investor restriction
	InvestorRestrictionAdd  endpoint.Endpoint
	InvestorRestrictionRead endpoint.Endpoint
	// securitisation
	EligibleAccount               endpoint.Endpoint
	EligibleAccountSummary        endpoint.Endpoint
	EligibleAccountSummaryReport  endpoint.Endpoint
	SecuritisationPoolAdd         endpoint.Endpoint
	SecuritisationPoolRead        endpoint.Endpoint
	SecuritisationPoolUpdate      endpoint.Endpoint
	SecuritisationPoolAll         endpoint.Endpoint
	SecuritisationDelete          endpoint.Endpoint
	SecuritisationDashboard       endpoint.Endpoint
	SecuritisationDashboardExport endpoint.Endpoint
	SecuritisationPoolReport      endpoint.Endpoint

	// product
	ProductAdd                  endpoint.Endpoint
	ProductSearch               endpoint.Endpoint
	ProductRead                 endpoint.Endpoint
	ProductUpdate               endpoint.Endpoint
	ProductDelete               endpoint.Endpoint
	ProductAnalysisCount        endpoint.Endpoint
	ProductOverpaymentLimit     endpoint.Endpoint
	ProductOverpaymentCalculate endpoint.Endpoint

	// Tasks
	TaskAdd        endpoint.Endpoint
	RecentTaskList endpoint.Endpoint
	TaskSummary    endpoint.Endpoint

	// lateFeeRule
	LateFeeRuleAdd    endpoint.Endpoint
	LateFeeRuleRead   endpoint.Endpoint
	LateFeeRuleUpdate endpoint.Endpoint
}

func NewEndpoints(
	lateFeeRuleSvc lateFeeRuleSvc.Service,
	accountSvc accountSvc.Service,
	duediligenceSvc duediligenceSvc.Service,
	accountflagSvc accountflagSvc.Service,
	accountAuditLogSvc accountAuditLogSvc.Service,
	accountLockRuleSvc accountLockRuleSvc.Service,
	serviceRestrictionSvc serviceRestrictionSvc.Service,
	investorRestrictionSvc investorRestrictionSvc.Service,
	dashboardSvc dashboardSvc.Service,
	securitisationSvc securitisationSvc.Service,
	tasksSvc tasksSvc.Service,

) Endpoints {
	return Endpoints{

		// Account
		AccountAdd:          makeAccountAddEndpoint(accountSvc),
		AccountRead:         makeAccountReadEndpoint(accountSvc),
		AccountUpdate:       makeAccountUpdateEndpoint(accountSvc),
		RecentArrears:       makeAccountRecentArrearsEndpoint(accountSvc),
		MortgagePerformance: makeAccountMortgagePerformanceEndpoint(accountSvc),

		// Due Diligence
		DueDiligenceAdd:       makeDueDiligenceAddEndpoint(duediligenceSvc),
		DueDiligenceRead:      makeDueDiligenceReadEndpoint(duediligenceSvc),
		DueDiligenceUpdate:    makeDueDiligenceUpdateEndpoint(duediligenceSvc),
		DueDiligenceByAccount: makeDueDiligenceByAccountEndpoint(duediligenceSvc),

		// account audit log
		AccountAuditLogRead: makeAccountAuditLogReadEndpoint(accountAuditLogSvc),

		// account flag
		AccountFlagAdd:  makeAccountFlagAddEndpoint(accountflagSvc),
		AccountFlagRead: makeAccountFlagReadEndpoint(accountflagSvc),

		// account lock rule
		AccountLockRuleAdd:  makeAccountLockRuleAddEndpoint(accountLockRuleSvc),
		AccountLockRuleRead: makeAccountLockRuleReadEndpoint(accountLockRuleSvc),

		// dashboard
		PortfolioSummary: makeDashboardEndpoint(dashboardSvc),
		// service restriction
		ServiceRestrictionAdd:           makeServiceRestrictionAddEndpoint(serviceRestrictionSvc),
		ServiceRestrictionRead:          makeServiceRestrictionReadEndpoint(serviceRestrictionSvc),
		ServiceRestrictionReadByAccount: makeServiceRestrictionReadByAccountEndpoint(serviceRestrictionSvc),

		// investor restriction
		InvestorRestrictionAdd:  makeInvestorRestrictionAddEndpoint(investorRestrictionSvc),
		InvestorRestrictionRead: makeInvestorRestrictionReadEndpoint(investorRestrictionSvc),

		// tasks
		TaskAdd:        makeTaskAddEndpoint(tasksSvc),
		RecentTaskList: makeRecentTaskEndpoint(tasksSvc),
		TaskSummary:    makeTaskSummaryEndpoint(tasksSvc),

		// securitisation
		EligibleAccount:               makeAccountEligibleEndpoint(securitisationSvc),
		EligibleAccountSummary:        MakeEligibleAccountSummaryHandler(securitisationSvc),
		EligibleAccountSummaryReport:  MakeEligibleAccountSummaryReportHandler(securitisationSvc),
		SecuritisationPoolAdd:         MakeSecuritisationPoolAddHandler(securitisationSvc),
		SecuritisationPoolRead:        MakeSecuritisationPoolReadHandler(securitisationSvc),
		SecuritisationPoolUpdate:      MakeSecuritisationPoolUpdateHandler(securitisationSvc),
		SecuritisationPoolAll:         MakeSecuritisationPoolAllHandler(securitisationSvc),
		SecuritisationDelete:          MakeSecuritisationDeleteHandler(securitisationSvc),
		SecuritisationDashboard:       MakeSecuritisationDashboardHandler(securitisationSvc),
		SecuritisationDashboardExport: MakeSecuritisationDashboardExportHandler(securitisationSvc),
		SecuritisationPoolReport:      MakeSecuritisationPoolReportHandler(securitisationSvc),

		LateFeeRuleAdd:    makeLateFeeRuleAddEndpoint(lateFeeRuleSvc),
		LateFeeRuleRead:   makeLateFeeRuleReadEndpoint(lateFeeRuleSvc),
		LateFeeRuleUpdate: makeLateFeeRuleUpdateEndpoint(lateFeeRuleSvc),
	}
}
