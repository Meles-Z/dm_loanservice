package customer

import "time"

type CustomerType string

const (
	Residential CustomerType = "Residential"
	Commercial  CustomerType = "Commercial"
	BTL         CustomerType = "BTL"
	Land        CustomerType = "Land"
)

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

type KYCCheckStatus string

const (
	Passed      KYCCheckStatus = "Passed"
	Failed      KYCCheckStatus = "Failed"
	UnderReview KYCCheckStatus = "UnderReview"
)

type RiskProfile string

const (
	High   RiskProfile = "High"
	Medium RiskProfile = "Medium"
	Low    RiskProfile = "Low"
)

type Customer struct {
	ID                       string         `json:"id"`
	CustomerType             CustomerType   `json:"customer_type"`
	FirstName                string         `json:"first_name"`
	LastName                 string         `json:"last_name"`
	DateOfBirth              time.Time      `json:"date_of_birth"`
	PlaceOfBirth             string         `json:"place_of_birth,omitempty"`
	Gender                   Gender         `json:"gender,omitempty"`
	Nationality              string         `json:"nationality,omitempty"`
	NationalInsuranceNo      string         `json:"national_insurance_no,omitempty"`
	Email                    string         `json:"email,omitempty"`
	Phone                    string         `json:"phone,omitempty"`
	Address                  string         `json:"address,omitempty"`
	RiskProfile              RiskProfile    `json:"risk_profile,omitempty"`
	FraudFlag                bool           `json:"fraud_flag"`
	FraudNotes               string         `json:"fraud_notes,omitempty"`
	KYCCheckStatus           KYCCheckStatus `json:"kyc_check_status,omitempty"`
	AMLFlag                  bool           `json:"aml_flag"`
	PoliticallyExposedPerson bool           `json:"politically_exposed_person"`
	VulnerabilityFlag        bool           `json:"vulnerability_flag"`
	ConsentFlag              bool           `json:"consent_flag"`
	ConsentWithdrawalDate    *time.Time     `json:"consent_withdrawal_date,omitempty"`
	CreatedAt                time.Time      `json:"created_at" db:"created_at"`
	CreatedBy                string         `json:"-" db:"created_by"`
	UpdatedAt                time.Time      `json:"updated_at" db:"updated_at"`
	UpdatedBy                string         `json:"-" db:"updated_by"`
	DeletedAt                *time.Time     `json:"deleted_at" db:"deleted_at"`
}

type Query struct {
	CustomerID string     `json:"customer_id,omitempty"` // exact match
	Name       string     `json:"name,omitempty"`        // partial match (first_name OR last_name)
	DOB        *time.Time `json:"dob,omitempty"`         // exact date match
	Email      string     `json:"email,omitempty"`       // exact match
	Phone      string     `json:"phone,omitempty"`       // exact match
	Address    string     `json:"address"`
	SortBy     string     `json:"sort_by,omitempty"`    // e.g., "id", "name", "date_of_birth"
	SortOrder  string     `json:"sort_order,omitempty"` // "asc" or "desc"
}
