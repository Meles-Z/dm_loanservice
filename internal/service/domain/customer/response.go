package customer

import (
	"dm_loanservice/drivers/dbmodels"
	"time"
)

type (
	CustomerSearchResponse struct {
		ID              string              `json:"id"`
		Name            string              `json:"name"`
		Email           string              `json:"email"`
		Phone           string              `json:"phone"`
		Location        string              `json:"location"`
		DateOfBirth     time.Time           `json:"date_of_birth"`
		ActiveMortgages int                 `json:"active_mortgages"`
		LastUpdated     time.Time           `json:"last_updated"`
		LinkedMortgage  []*dbmodels.Product `json:"linked_mortgage"`
	}
)

type CustomerListResponse struct {
	Data       []CustomerSearchResponse `json:"customers"`
	Pagination CustomerPagination       `json:"pagination"`
}

type CustomerReadResponse struct {
	Customer *Customer `json:"products"`
}

type CustomerUpdateResponse struct {
	Customer *Customer `json:"products"`
}

type CustomerDeleteResponse struct {
	Success bool `json:"success"`
}

type CustomerPagination struct {
	Page      int32 `json:"page"`
	Length    int32 `json:"length"`
	Total     int32 `json:"total"`
	TotalPage int32 `json:"total_page"`
}
