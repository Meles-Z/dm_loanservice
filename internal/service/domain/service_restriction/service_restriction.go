package servicerestriction

type ServiceRestriction struct {
	ID              string `json:"id"`
	AccountID       string `json:"account_id"`
	FieldName       string `json:"field_name"`
	RestrictionType string `json:"restriction_type"`
	ActionName      string `json:"action_name"`
	IsActive        bool   `json:"is_active"`
	Reason          string `json:"reason"`
	Source          string `json:"source"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}
