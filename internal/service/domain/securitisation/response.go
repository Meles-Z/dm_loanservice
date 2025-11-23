package securitisation

type EligibleAccountResponse struct {
	Page         int                `json:"page"`
	PageSize     int                `json:"page_size"`
	TotalRecords int                `json:"total_records"`
	Loans        []EligibleLoanItem `json:"loans"`
}

type EligibleLoanItem struct {
	LoanID            string       `json:"loan_id"`
	AccountNumber     string       `json:"account_number"`
	Borrower          BorrowerInfo `json:"borrower"`
	Property          PropertyInfo `json:"property"`
	LTV               float64      `json:"ltv"`
	ArrearsDays       int          `json:"arrears_days"`
	OriginationDate   string       `json:"origination_date"`
	DDStatus          string       `json:"dd_status"`
	EligibilityStatus string       `json:"eligibility_status"`
	Flags             []string     `json:"flags"`
}

type BorrowerInfo struct {
	FullName string `json:"full_name"`
	Region   string `json:"region"`
}

type PropertyInfo struct {
	PropertyType string  `json:"property_type"`
	Region       string  `json:"region"`
	Valuation    float64 `json:"valuation"`
}

type EligibleAccountSummaryResponse struct {
	TotalLoans          int                 `json:"total_loans"`
	EligibleLoans       int                 `json:"eligible_loans"`
	IneligibleLoans     int                 `json:"ineligible_loans"`
	DueDiligenceSummary DueDiligenceSummary `json:"due_diligence"`
	FlagSummary         FlagSummary         `json:"flags"`
}

type DueDiligenceSummary struct {
	Pass    int `json:"pass"`
	Fail    int `json:"fail"`
	Pending int `json:"pending"`
}

type FlagSummary struct {
	SecReady     int `json:"sec_ready"`
	SecExcluded  int `json:"sec_excluded"`
	ManualReview int `json:"manual_review"`
}

type EligibleAccountSummaryReportResponse struct {
    File     string `json:"file"`
    FileName string `json:"file_name"`
}