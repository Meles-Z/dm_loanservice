package accountflag

import "dm_loanservice/drivers/validator"

type (
	AccountFlagAddRequest struct {
		AccountID string `json:"account_id" validate:"required"`
		FlagType  string `json:"flag_type" validate:"required"`
		Reason    string `json:"reason" validate:"required"`
	}

	AccountFlagReadRequest struct {
		ID string `json:"id" validate:"required"`
	}
)

func (r *AccountFlagAddRequest) Validate() error {
	return validator.Validate.Struct(&r)
}
func (r *AccountFlagReadRequest) Validate() error {
	return validator.Validate.Struct(&r)
}
