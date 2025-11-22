
package accountauditlog

import "dm_loanservice/drivers/validator"

type AccountAuditLogAddRequest struct {
	AccountID   string `json:"account_id" validate:"required"`
	Action      string `json:"action" validate:"required"`
	Details     string `json:"details" validate:"required"`
	PerformedBy int    `json:"performed_by" validate:"required"`
}

type AccountAuditLogReadRequest struct {
	ID string `json:"id" validate:"required"`
}

func (r *AccountAuditLogAddRequest) Validate() error {
	return validator.Validate.Struct(r)
}
func (r *AccountAuditLogReadRequest) Validate() error {
	return validator.Validate.Struct(r)
}

