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

type SecuritisationPoolAddResponse struct {
	Pool *SecuritisationPool `json:"pool"`
}

type SecuritisationPoolReadResponse struct {
	Pool *SecuritisationPool `json:"pool"`
}

type SecuritisationPoolUpdateResponse struct {
	Pool *SecuritisationPool `json:"pool"`
}

type SecuritisationPoolAllResponse struct {
	Pools []*SecuritisationPool `json:"pools"`
}

type DashboardEligibleAccountResponse struct {
	TotalLoans          int64               `json:"total_loans"`
	EligibleLoans       int64               `json:"eligible_loans"`
	IneligibleLoans     int64               `json:"ineligible_loans"`
	TotalOutstanding    float64             `json:"total_outstanding_balance"`
	DueDiligenceSummary DueDiligenceSummary `json:"due_diligence"`
	FlagSummary         FlagSummary         `json:"flags"`
	// DailyChanges        DashboardDailyChanges `json:"daily_changes"`
	// Snapshot            DashboardSnapshot     `json:"snapshot"`
	// Trends              DashboardTrends       `json:"trends"`
}

type DailyChanges struct {
	NewLoansToday     int64 `json:"new_loans_today"`
	LoansFlaggedToday int64 `json:"loans_flagged_today"`
}

type DashboardExportResponse struct {
	FileName string `json:"file_name"`
	MimeType string `json:"mime_type"`
	FileData string `json:"file"` // base64
}

/*
{
  "snapshot": {
    "total_loans": 1200,
    "eligible_loans": 850,
    "ineligible_loans": 350,

    "total_outstanding_balance": "245,000,000.00",
    "average_ltv": 63.4,
    "weighted_average_ltv": 61.9,
    "portfolio_arrears_rate": 4.3,

    "dd_summary": {
      "pass": 900,
      "fail": 150,
      "pending": 150
    },

    "flag_summary": {
      "sec_ready": 800,
      "sec_excluded": 200,
      "manual_review": 50
    },

    "daily_changes": {
      "new_loans_today": 12,
      "loans_flagged_today": 5
    }
  },

  "trends": {
    "eligibility": [
      { "date": "2025-01-01", "eligible": 820, "ineligible": 380 },
      { "date": "2025-01-02", "eligible": 830, "ineligible": 370 }
    ],
    "flags": [
      { "date": "2025-01-01", "sec_ready": 790, "sec_excluded": 220, "manual_review": 40 }
    ],
    "dd": [
      { "date": "2025-01-01", "pass": 890, "fail": 160, "pending": 150 }
    ],
    "balance": [
      { "date": "2025-01-01", "balance": 245000000.00 }
    ]
  }
}

*/

type SecuritisationPoolReportResponse struct {
	PoolID             string `json:"pool_id"`
	SPVName            string `json:"spv_name"`
	FundingSource      string `json:"funding_source"`
	PoolAllocationDate string `json:"pool_allocation_date"`
	LoanTransferDate   string `json:"loan_transfer_date"`
	NoteClass          string `json:"note_class"`
	ReportingCurrency  string `json:"reporting_currency"`
	//
	TotalLoans           int64   `json:"total_loans"`
	EligibleLoans        int64   `json:"eligible_loans"`
	IneligibleLoans      int64   `json:"ineligible_loans"`
	TotalOutstanding     float64 `json:"total_outstanding"`
	AverageLTV           float64 `json:"average_ltv"`
	WeightedAverageLTV   float64 `json:"weighted_average_ltv"`
	PortfolioArrearsRate float64 `json:"portfolio_arrears_rate"`
	//
	DDSummary   DueDiligenceSummary `json:"dd_summary"`
	FlagSummary FlagSummary         `json:"flag_summary"`
}

/*
{
  "pool_id": "123",
  "spv_name": "ABC SPV",
  "funding_source": "Bank X",
  "pool_allocation_date": "2025-11-01",
  "loan_transfer_date": "2025-11-02",
  "current_pool_balance": 45000000.50,
  "factor": 0.95,
  "note_class": "A",
  "reporting_currency": "USD",

  "total_loans": 120,
  "eligible_loans": 100,
  "ineligible_loans": 20,
  "total_outstanding": 45000000.50,

  "average_ltv": 63.4,
  "weighted_average_ltv": 61.9,
  "portfolio_arrears_rate": 4.3,

  "dd_summary": {
    "pass": 90,
    "fail": 20,
    "pending": 10
  },
  "flag_summary": {
    "sec_ready": 80,
    "sec_excluded": 15,
    "manual_review": 5
  },
  "audit_log": [
    {
      "report_id": "rpt123",
      "generated_by": "user1",
      "generated_at": "2025-11-15T10:30:00Z",
      "notes": "Initial compliance report"
    }
  ]
}


*/
