package tasks

import (
	"dm_loanservice/drivers/validator"
	"time"
)

type TaskAddRequest struct {
	InquiriesID string    `json:"inquiries_id" validate:"required"`
	CustomerID  string    `json:"customer_id" validate:"required"`
	AssignedTo  int       `json:"assigned_to" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	Status      string    `json:"status" validate:"required"`
	DueDate     time.Time `json:"due_date" validate:"required"`
	Priority    string    `json:"priority" validate:"required"`
	CreatedBy   string    `json:"created_by"`
}

func (a TaskAddRequest) Validate() error {
	return validator.Validate.Struct(a)
}
