package securitisation

type EligibleAccountRequest struct {
	ID string `json:"id"`
}

type EligibleLoansQuery struct {
	Page            int     `json:"page"`
	PageSize        int     `json:"page_size"`
	MortgageType    string  `json:"mortgage_type"`
	Region          string  `json:"region"`
	LTVMin          float64 `json:"ltv_min"`
	LTVMax          float64 `json:"ltv_max"`
	ArrearsDaysMax  int     `json:"arrears_days_max"`
	OriginationFrom string  `json:"origination_date_from"`
	OriginationTo   string  `json:"origination_date_to"`
	PropertyType    string  `json:"property_type"`
	SortBy          string  `json:"sort_by"`
	SortDirection   string  `json:"sort_direction"`
}
