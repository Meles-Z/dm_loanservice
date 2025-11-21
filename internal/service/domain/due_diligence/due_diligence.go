package duediligence

import "github.com/aarondl/null/v8"

type DueDiligence struct {
	ID            string      `json:"id"`
	AccountID     string      `json:"account_id"`
	ChecklistItem string      `json:"checklist_item"`
	Status        string      `json:"status"`
	Comments      null.String `json:"comments"`
	CreatedBy     null.Int    `json:"created_by"`
	CreatedAt     null.Time   `json:"created_at"`
	UpdatedAt     null.Time   `json:"updated_at"`
}
