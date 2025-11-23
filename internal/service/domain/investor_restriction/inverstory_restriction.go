package investorrestriction

type InvestorRestriction struct {
	ID               string `json:"id"`
	AccountID        string `json:"account_id"`
	RestrictionScope string `json:"restriction_scope"`
	FieldName        string `json:"field_name"`
	ActionName       string `json:"action_name"`
	RuleType         string `json:"rule_type"`
	Reason           string `json:"reason"`
	IsActive         bool   `json:"is_active"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}
