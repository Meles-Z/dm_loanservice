package accountlockrule

import "dm_loanservice/drivers/validator"

type AccountLockRuleAddRequest struct {
	AccountID  string `json:"account_id" validate:"required"`
	Status     string `json:"loan_status" validate:"required"`
	FieldName  string `json:"field_name" validate:"required"`
	LockType   string `json:"lock_type" validate:"required"`
	LockReason string `json:"lock_reason" validate:"required"`
}

type AccountLockRuleReadRequest struct {
	ID string `json:"id" validate:"required"`
}

func (r *AccountLockRuleAddRequest) Validate() error {
	return validator.Validate.Struct(r)
}

func (r *AccountLockRuleReadRequest) Validate() error {
	return validator.Validate.Struct(r)
}
