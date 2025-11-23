package property

type Property struct {
	ID           string `json:"id"`
	Address      string `json:"address"`
	PropertyType string `json:"property_type"`
	Region       string `json:"region"`
	Valuation    string `json:"valuation"`
	SizeSQFT     string `json:"size_sqft"`
	YearBuilt    string `json:"year_built"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
