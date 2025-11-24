package securitisation

import "github.com/aarondl/sqlboiler/v4/types"

type SecuritisationPool struct {
	ID                       string        `json:"id"`
	FundingSource            string        `json:"funding_source"`
	ServicingRole            string        `json:"servicing_role"`
	SPVName                  string        `json:"spv_name"`
	SPVJurisdiction          string        `json:"spv_jurisdiction"`
	PoolAllocationDate       string        `json:"pool_allocation_date"`
	LoanTransferDate         string        `json:"loan_transfer_date"`
	CurrentPoolBalance       types.Decimal `json:"current_pool_balance"`
	Factor                   types.Decimal `json:"factor"`
	NoteClass                string        `json:"note_class"`
	InterestRemittance       string        `json:"interest_remittance_date"`
	PrincipalRemittance      string        `json:"principal_remittance_date"`
	ServicingFeeRate         types.Decimal `json:"servicing_fee_rate"`
	ReportingCurrency        string        `json:"reporting_currency"`
	ESMAAssetCode            string        `json:"esma_asset_code"`
	CreditEnhancementType    string        `json:"credit_enhancement_type"`
	InvestorReportIdentifier string        `json:"investor_report_identifier"`
	CreatedAt                string        `json:"created_at"`
	UpdatedAt                string        `json:"updated_at"`
}
