package collateral

type Collateral struct {
	ID             string `json:"id"`
	AccountID      string `json:"account_id"`
	PropertyID     string `json:"property_id"`
	CollateralType string `json:"collateral_type"`
	SecurityType   string `json:"security_type"`
	LienPosition   string `json:"lien_position"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
