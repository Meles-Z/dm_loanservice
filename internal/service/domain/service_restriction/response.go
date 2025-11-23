package servicerestriction

type ServiceRestrictionResponse struct {
	ServiceRestriction *ServiceRestriction `json:"service_restriction"`
}

type ServiceRestrictionSliceResponse struct {
	OverallRestrictionStatus string            `json:"overall_restriction_status"`
	Restrictions             []RestrictionItem `json:"restrictions"`
}

type RestrictionItem struct {
	ID               string `json:"id"`
	Type             string `json:"type"` // loan_lock_rule | servicing_restriction | investor_restriction
	RestrictionLevel string `json:"restriction_level"`

	FieldName  *string `json:"field_name,omitempty"`
	ActionName *string `json:"action_name,omitempty"`

	Reason string `json:"reason"`
	Source string `json:"source"`

	InvestorID *string `json:"investor_id,omitempty"`
}
