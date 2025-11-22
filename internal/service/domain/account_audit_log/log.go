package accountauditlog

import "time"

type AccountAuditLog struct {
	ID          string    `json:"id"`
	AccountID   string    `json:"account_id"`
	Action      string    `json:"action"`
	Details     string    `json:"details"`
	PerformedBy int       `json:"performed_by"`
	PerformedAt time.Time `json:"performed_at"`
}
