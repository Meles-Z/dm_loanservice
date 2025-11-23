package collateral

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
)

type Repository interface {
	CollateralAdd(ctx context.Context, m dbmodels.Collateral) (*dbmodels.Collateral, error)
	CollateralRead(ctx context.Context, collateralID string) (*dbmodels.Collateral, error)
}
