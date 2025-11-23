package property

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
)

type Repository interface {
	PropertyAdd(ctx context.Context, m dbmodels.Property) (*dbmodels.Property, error)
	PropertyRead(ctx context.Context, propertyID string) (*dbmodels.Property, error)
}
