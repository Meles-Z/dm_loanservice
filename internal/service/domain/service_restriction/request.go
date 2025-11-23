package servicerestriction

import "dm_loanservice/drivers/validator"

type ServiceRestrictionAddRequest struct {
	ID              string `json:"id" validate:"required"`
	AccountID       string `json:"account_id" validate:"required"`
	FieldName       string `json:"field_name" validate:"required"`
	RestrictionType string `json:"restriction_type" validate:"required"`
	ActionName      string `json:"action_name" validate:"required"`
	IsActive        bool   `json:"is_active" validate:"required"`
	Reason          string `json:"reason" validate:"required"`
	Source          string `json:"source" validate:"required"`
}

type ServiceRestrictionReadRequest struct {
	ID string `json:"id" validate:"required"`
}

type ServiceRestrictionReadByAccountRequest struct {
	ID string `json:"id" validate:"required"`
}

func (r *ServiceRestrictionAddRequest) Validate() error {
	return validator.Validate.Struct(r)
}

func (r *ServiceRestrictionReadRequest) Validate() error {
	return validator.Validate.Struct(r)
}

func (r *ServiceRestrictionReadByAccountRequest) Validate() error {
	return validator.Validate.Struct(r)
}
