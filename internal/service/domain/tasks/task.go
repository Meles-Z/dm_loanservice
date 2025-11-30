package tasks

import "time"

type Tasks struct {
	ID          string     `json:"id"`
	AssignedTo  int        `json:"assigned_to"`
	InquiriesID string     `json:"inquiries_id"`
	CustomerID  string     `json:"customer_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	DueDate     time.Time  `json:"due_date"`
	Priority    string     `json:"priority"`
	CreatedBy   string     `json:"created_by"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
