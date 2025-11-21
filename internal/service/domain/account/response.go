package account

import (
	"time"
)

type AccountResponse struct {
	ID         string    `json:"id"`
	MortgageID string    `json:"mortgage_id"`
	CustomerID string    `json:"customer_id"`
	ProductID  string    `json:"product_id"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
}

type AccountRecentResponse struct {
	CustomerName string  `json:"customer_name"`
	MortgageID   string  `json:"mortgage_id"`
	ArrearsAge   int     `json:"arrears_age"` // e.g., days
	Amount       float64 `json:"amount"`
	Status       string  `json:"status"`
}

type AccountReadResponse struct {
	Account Accounts `json:"account"`
}