package account

import (
	"time"

	"github.com/aarondl/sqlboiler/v4/types"
)

type ForbearanceType string

const (
	ForbearancePaymentHoliday ForbearanceType = "Payment Holiday"
	ForbearanceInterestOnly   ForbearanceType = "Interest Only"
	ForbearanceReducedPayment ForbearanceType = "Reduced Payment"
)

type CollateralType string

const (
	CollateralHouse  CollateralType = "House"
	CollateralFlat   CollateralType = "Flat"
	CollateralOffice CollateralType = "Office"
	CollateralLand   CollateralType = "Land"
)

type SecurityType string

const (
	SecurityMortgageDeed SecurityType = "Mortgage Deed"
	SecurityLegalCharge  SecurityType = "Legal Charge"
	SecurityGuarantee    SecurityType = "Guarantee"
)

type StressTestResult string

const (
	StressPass StressTestResult = "Pass"
	StressFail StressTestResult = "Fail"
)

type Accounts struct {
	ID                 string        `json:"id"`
	MortgageID         string        `json:"mortgage_id" db:"mortgage_id"`
	CustomerID         string        `json:"customer_id" db:"customer_id"`
	ProductID          string        `json:"product_id" db:"product_id"` // FK → MortgageProduct
	LoanAmount         types.Decimal `json:"loan_amount" db:"loan_amount"`
	BalanceOutstanding types.Decimal `json:"balance_outstanding" db:"balance_outstanding"`
	StartDate          time.Time     `json:"start_date" db:"start_date"`
	EndDate            time.Time     `json:"end_date" db:"end_date"`
	TermYears          int           `json:"term_years" db:"term_years"`

	ArrearsFlag   bool          `json:"arrears_flag" db:"arrears_flag"`
	ArrearsAmount types.Decimal `json:"arrears_amount" db:"arrears_amount"`
	ArrearsDays   int           `json:"arrears_days" db:"arrears_days"`

	ForbearanceFlag bool            `json:"forbearance_flag" db:"forbearance_flag"`
	ForbearanceType ForbearanceType `json:"forbearance_type" db:"forbearance_type"`

	FraudFlag  bool   `json:"fraud_flag" db:"fraud_flag"`
	FraudNotes string `json:"fraud_notes" db:"fraud_notes"`

	RedrawFacility    bool           `json:"redraw_facility" db:"redraw_facility"`
	CollateralAddress string         `json:"collateral_address" db:"collateral_address"`
	CollateralType    CollateralType `json:"collateral_type" db:"collateral_type"`
	SecurityType      SecurityType   `json:"security_type" db:"security_type"`

	PortfolioID         string           `json:"portfolio_id" db:"portfolio_id"`
	StressTestResult    StressTestResult `json:"stress_test_result" db:"stress_test_result"`
	CapitalAdequacyFlag bool             `json:"capital_adequacy_flag" db:"capital_adequacy_flag"`

	// Audit fields
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	CreatedBy string     `json:"-" db:"created_by"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	UpdatedBy string     `json:"-" db:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

