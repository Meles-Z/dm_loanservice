package accountflag

import "time"

type AccountFlag struct {
	ID        string    `json:"id"`
	AccountID string    `json:"account_id"`
	FlagType  string    `json:"flag_type"`
	Reason    string    `json:"reason"`
	FlaggedBy int       `json:"flagged_by"`
	FlaggedAt time.Time `json:"flagged_at"`
}
