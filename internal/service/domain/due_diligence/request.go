package duediligence

import "dm_loanservice/drivers/validator"

type (
	DueDiligenceAddRequest struct {
		AccountID     string `json:"account_id" validate:"required"`
		ChecklistItem string `json:"checklist_item" validate:"required"`
		Status        string `json:"status" validate:"required"`
		Comments      string `json:"comments" validate:"required"`
	}

	DueDiligenceReadRequest struct {
		ID string `json:"id" validate:"required"`
	}

	DueDiligenceUpdateRequest struct {
		ID            string  `json:"id" validate:"required"`
		ChecklistItem *string `json:"checklist_item"`
		Status        *string `json:"status"`
		Comments      *string `json:"comments"`
	}

	DueDiligenceByAccountRequest struct {
		AccountID string `json:"account_id" validate:"required"`
	}
)

func (r *DueDiligenceAddRequest) Validate() error {
	return validator.Validate.Struct(&r)
}
func (r *DueDiligenceReadRequest) Validate() error {
	return validator.Validate.Struct(&r)
}
func (r *DueDiligenceUpdateRequest) Validate() error {
	return validator.Validate.Struct(&r)
}
func (r *DueDiligenceByAccountRequest) Validate() error {
	return validator.Validate.Struct(&r)
}
