package securitisation

import (
	"dm_loanservice/drivers/validator"
	"time"

	"github.com/aarondl/sqlboiler/v4/types"
)

type EligibleAccountRequest struct {
	ID string `json:"id"`
}

type EligibleLoansQuery struct {
	Page            int     `json:"page"`
	PageSize        int     `json:"page_size"`
	MortgageType    string  `json:"mortgage_type"`
	Region          string  `json:"region"`
	LTVMin          float64 `json:"ltv_min"`
	LTVMax          float64 `json:"ltv_max"`
	ArrearsDaysMax  int     `json:"arrears_days_max"`
	OriginationFrom string  `json:"origination_date_from"`
	OriginationTo   string  `json:"origination_date_to"`
	PropertyType    string  `json:"property_type"`
	SortBy          string  `json:"sort_by"`
	SortDirection   string  `json:"sort_direction"`
}

type SecuritisationPoolAddRequest struct {
	FundingSource            string        `json:"funding_source"`
	ServicingRole            string        `json:"servicing_role"`
	SPVName                  string        `json:"spv_name"`
	SPVJurisdiction          string        `json:"spv_jurisdiction"`
	PoolAllocationDate       time.Time     `json:"pool_allocation_date"`
	LoanTransferDate         time.Time     `json:"loan_transfer_date"`
	CurrentPoolBalance       types.Decimal `json:"current_pool_balance"`
	Factor                   types.Decimal `json:"factor"`
	NoteClass                string        `json:"note_class"`
	InterestRemittance       time.Time     `json:"interest_remittance_date"`
	PrincipalRemittance      time.Time     `json:"principal_remittance_date"`
	ServicingFeeRate         types.Decimal `json:"servicing_fee_rate"`
	ReportingCurrency        string        `json:"reporting_currency"`
	ESMAAssetCode            string        `json:"esma_asset_code"`
	CreditEnhancementType    string        `json:"credit_enhancement_type"`
	InvestorReportIdentifier string        `json:"investor_report_identifier"`
}

type SecuritisationPoolReadRequest struct {
	ID string `json:"id"`
}

type SecuritisationPoolUpdateRequest struct {
	ID                       string         `json:"id"`
	FundingSource            *string        `json:"funding_source,omitempty"`
	ServicingRole            *string        `json:"servicing_role,omitempty"`
	SPVName                  *string        `json:"spv_name,omitempty"`
	SPVJurisdiction          *string        `json:"spv_jurisdiction,omitempty"`
	PoolAllocationDate       *string        `json:"pool_allocation_date,omitempty"`
	LoanTransferDate         *string        `json:"loan_transfer_date,omitempty"`
	CurrentPoolBalance       *types.Decimal `json:"current_pool_balance,omitempty"`
	Factor                   *types.Decimal `json:"factor,omitempty"`
	NoteClass                *string        `json:"note_class,omitempty"`
	InterestRemittance       *string        `json:"interest_remittance_date,omitempty"`
	PrincipalRemittance      *string        `json:"principal_remittance_date,omitempty"`
	ServicingFeeRate         *types.Decimal `json:"servicing_fee_rate,omitempty"`
	ReportingCurrency        *string        `json:"reporting_currency,omitempty"`
	ESMAAssetCode            *string        `json:"esma_asset_code,omitempty"`
	CreditEnhancementType    *string        `json:"credit_enhancement_type,omitempty"`
	InvestorReportIdentifier *string        `json:"investor_report_identifier,omitempty"`
}

type SecuritisationDeleteRequest struct {
	ID string `json:"id"`
}

type DashboardExportRequest struct {
	Format string `json:"format"` // csv or xlsx
}

type SecuritisationPoolReportRequest struct {
	ID string `json:"id"`
}

func (r *SecuritisationPoolAddRequest) Validate() error {
	return validator.Validate.Struct(r)
}

func (r *SecuritisationPoolReadRequest) Validate() error {
	return validator.Validate.Struct(r)
}

func (r *SecuritisationPoolUpdateRequest) Validate() error {
	return validator.Validate.Struct(r)
}
func (r *SecuritisationDeleteRequest) Validate() error {
	return validator.Validate.Struct(r)
}

func (r *SecuritisationPoolReportRequest) Validate() error {
	return validator.Validate.Struct(r)
}
