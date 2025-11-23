package account

import (
	"time"

	"github.com/aarondl/sqlboiler/v4/types"
)

type AccountAddRequest struct {
	MortgageID         string        `json:"mortgage_id"`
	CustomerID         string        `json:"customer_id" validate:"required"`
	ProductID          string        `json:"product_id" validate:"required"`
	LoanAmount         types.Decimal `json:"loan_amount" validate:"required"`
	BalanceOutstanding types.Decimal `json:"balance_outstanding" validate:"required"`
	StartDate          time.Time     `json:"start_date"`
	EndDate            time.Time     `json:"end_date"`
	TermYears          int           `json:"term_years"`

	ArrearsFlag   bool          `json:"arrears_flag"`
	ArrearsAmount types.Decimal `json:"arrears_amount"`
	ArrearsDays   int           `json:"arrears_days"`

	ForbearanceFlag bool            `json:"forbearance_flag"`
	ForbearanceType ForbearanceType `json:"forbearance_type"`

	FraudFlag  bool   `json:"fraud_flag"`
	FraudNotes string `json:"fraud_notes"`

	RedrawFacility    bool           `json:"redraw_facility"`
	CollateralAddress string         `json:"collateral_address"`
	CollateralType    CollateralType `json:"collateral_type"`
	SecurityType      SecurityType   `json:"security_type"`

	PortfolioID         string           `json:"portfolio_id"`
	StressTestResult    StressTestResult `json:"stress_test_result"`
	CapitalAdequacyFlag bool             `json:"capital_adequacy_flag"`
}

type XMLAccount struct {
	MortgageID         string  `xml:"MortgageID"`
	LoanAmount         float64 `xml:"LoanAmount"`
	BalanceOutstanding float64 `xml:"BalanceOutstanding"`
	StartDate          string  `xml:"StartDate"`
	EndDate            string  `xml:"EndDate"`
	TermYears          int     `xml:"TermYears"`

	ArrearsFlag   bool    `xml:"ArrearsFlag"`
	ArrearsAmount float64 `xml:"ArrearsAmount"`
	ArrearsDays   int     `xml:"ArrearsDays"`

	ForbearanceFlag bool   `xml:"ForbearanceFlag"`
	ForbearanceType string `xml:"ForbearanceType"`

	FraudFlag  bool   `xml:"FraudFlag"`
	FraudNotes string `xml:"FraudNotes"`

	RedrawFacility    bool   `xml:"RedrawFacility"`
	CollateralAddress string `xml:"CollateralAddress"`
	CollateralType    string `xml:"CollateralType"`
	SecurityType      string `xml:"SecurityType"`

	PortfolioID         string `xml:"PortfolioID"`
	StressTestResult    string `xml:"StressTestResult"`
	CapitalAdequacyFlag bool   `xml:"CapitalAdequacyFlag"`
}

type JSONAccount struct {
	MortgageID         string  `json:"mortgage_id"`
	LoanAmount         float64 `json:"loan_amount"`
	BalanceOutstanding float64 `json:"balance_outstanding"`
	StartDate          string  `json:"start_date"`
	EndDate            string  `json:"end_date"`
	TermYears          int     `json:"term_years"`

	ArrearsFlag   bool    `json:"arrears_flag"`
	ArrearsAmount float64 `json:"arrears_amount"`
	ArrearsDays   int     `json:"arrears_days"`

	ForbearanceFlag bool   `json:"forbearance_flag"`
	ForbearanceType string `json:"forbearance_type"`

	FraudFlag  bool   `json:"fraud_flag"`
	FraudNotes string `json:"fraud_notes"`

	RedrawFacility    bool   `json:"redraw_facility"`
	CollateralAddress string `json:"collateral_address"`
	CollateralType    string `json:"collateral_type"`
	SecurityType      string `json:"security_type"`

	PortfolioID         string `json:"portfolio_id"`
	StressTestResult    string `json:"stress_test_result"`
	CapitalAdequacyFlag bool   `json:"capital_adequacy_flag"`
}

type AccountReadRequest struct {
	ID string `json:"account" validate:"required"`
}

type AccountUpdateRequest struct {
	ID                 string         `json:"id" validate:"required"`
	MortgageID         *string        `json:"mortgage_id,omitempty"`
	CustomerID         *string        `json:"customer_id,omitempty"`
	ProductID          *string        `json:"product_id,omitempty"`
	LoanAmount         *types.Decimal `json:"loan_amount,omitempty"`
	BalanceOutstanding *types.Decimal `json:"balance_outstanding,omitempty"`
	StartDate          *time.Time     `json:"start_date,omitempty"`
	EndDate            *time.Time     `json:"end_date,omitempty"`
	TermYears          *int           `json:"term_years,omitempty"`

	ArrearsFlag   *bool          `json:"arrears_flag,omitempty"`
	ArrearsAmount *types.Decimal `json:"arrears_amount,omitempty"`
	ArrearsDays   *int           `json:"arrears_days,omitempty"`

	ForbearanceFlag *bool   `json:"forbearance_flag,omitempty"`
	ForbearanceType *string `json:"forbearance_type,omitempty"`

	FraudFlag  *bool   `json:"fraud_flag,omitempty"`
	FraudNotes *string `json:"fraud_notes,omitempty"`

	RedrawFacility    *bool   `json:"redraw_facility,omitempty"`
	CollateralAddress *string `json:"collateral_address,omitempty"`
	CollateralType    *string `json:"collateral_type,omitempty"`
	SecurityType      *string `json:"security_type,omitempty"`

	PortfolioID         *string `json:"portfolio_id,omitempty"`
	StressTestResult    *string `json:"stress_test_result,omitempty"`
	CapitalAdequacyFlag *bool   `json:"capital_adequacy_flag,omitempty"`
}

func (t *AccountReadRequest) Validate() error {
	return nil
}
