package customer

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
)

type Repository interface {
	CustomerAdd(ctx context.Context, m dbmodels.Customer) (*dbmodels.Customer, error)
	SearchCustomers(ctx context.Context, page, length int, params Query) (*dbmodels.CustomerSlice, int64, error)
	FindCustomerById(ctx context.Context, id string) (*dbmodels.Customer, error)
	FindCustomersByProductID(ctx context.Context, productID string) ([]*dbmodels.Customer, error)
	UpdateCustomer(ctx context.Context, m dbmodels.Customer) (*dbmodels.Customer, error)
	DeleteCustomer(ctx context.Context, id string) error
}
