package investorrestriction

import "dm_loanservice/drivers/validator"

type InvestorRestrictionAddRequest struct {
	ID               string `json:"id" validate:"required"`
	AccountID        string `json:"account_id" validate:"required"`
	RestrictionScope string `json:"restriction_scope" validate:"required"`
	FieldName        string `json:"field_name" validate:"required"`
	ActionName       string `json:"action_name" validate:"required"`
	RuleType         string `json:"rule_type" validate:"required"`
	Reason           string `json:"reason" validate:"required"`
	IsActive         bool   `json:"is_active" validate:"required"`
}


type InvestorRestrictionReadRequest struct {
	ID string `json:"id" validate:"required"`
}

func (r *InvestorRestrictionAddRequest) Validate() error {
	return validator.Validate.Struct(r)
}

func (r *InvestorRestrictionReadRequest) Validate() error {
	return validator.Validate.Struct(r)
}
