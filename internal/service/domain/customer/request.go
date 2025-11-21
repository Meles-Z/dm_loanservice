package customer

import (
	"dm_loanservice/drivers/validator"
	"time"
)

type CustomerRequest struct {
	CustomerType             CustomerType   `json:"customer_type" validate:"required,oneof=Residential Commercial BTL Land"`
	FirstName                string         `json:"first_name" validate:"required"`
	LastName                 string         `json:"last_name" validate:"required"`
	DateOfBirth              time.Time      `json:"date_of_birth" validate:"required"`
	PlaceOfBirth             string         `json:"place_of_birth" validate:"required"`
	Gender                   Gender         `json:"gender" validate:"required"`
	Nationality              string         `json:"nationality" validate:"required"`
	NationalInsuranceNo      string         `json:"national_insurance_no" validate:"required"`
	Email                    string         `json:"email" validate:"required"`
	Phone                    string         `json:"phone" validate:"required"`
	Address                  string         `json:"address" validate:"required"`
	RiskProfile              RiskProfile    `json:"risk_profile" validate:"required,oneof=High Medium Low"`
	KYCCheckStatus           KYCCheckStatus `json:"kyc_check_status" validate:"required,oneof=Passed Failed UnderReview"`
	FraudFlag                bool           `json:"fraud_flag,omitempty"`
	FraudNotes               string         `json:"fraud_notes,omitempty"`
	AMLFlag                  bool           `json:"aml_flag,omitempty"`
	PoliticallyExposedPerson bool           `json:"politically_exposed_person,omitempty" validate:"required"`
	VulnerabilityFlag        bool           `json:"vulnerability_flag,omitempty"`
	ConsentFlag              bool           `json:"consent_flag,omitempty"`
	ConsentWithdrawalDate    *time.Time     `json:"consent_withdrawal_date" validate:"required"`
}

type CustomerSearchRequest struct {
	Page           int32  `json:"page,omitempty" validate:"gt=0"`
	Length         int32  `json:"length,omitempty" validate:"gt=0"`
	SearchKey      string `json:"search_key,omitempty"`
	SearchValue    string `json:"search_value,omitempty"`
	OrderKey       string `json:"order_key,omitempty"`
	OrderDirection string `json:"order_direction,omitempty"`
}

type CustomerDetailReadRequest struct {
	ID string `json:"id" validate:"required"`
}

type CustomerUpdateRequest struct {
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
}
type CustomerDeleteRequest struct {
	ID string `json:"id" validate:"required"`
}

func (c CustomerSearchRequest) Validate() error {
	if err := validator.Validate.Struct(c); err != nil {
		return err
	}

	if c.SearchValue == "" {
		if c.Length == 0 {
			c.Length = 25
		}
		if c.Page == 0 {
			c.Page = 1
		}
	}

	return nil
}

func (c CustomerRequest) Validate() error {
	return validator.Validate.Struct(c)
}

func (c CustomerDetailReadRequest) Validate() error {
	return validator.Validate.Struct(c)
}

func (c CustomerUpdateRequest) Validate() error {
	return validator.Validate.Struct(c)
}

func (c CustomerDeleteRequest) Validate() error {
	return validator.Validate.Struct(c)
}
