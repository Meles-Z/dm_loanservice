package securitisation

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
)

type Repository interface {
	SecuritisationPoolAdd(ctx context.Context, m dbmodels.SecuritisationPool) (*dbmodels.SecuritisationPool, error)
	SecuritisationPoolRead(ctx context.Context, poolID string) (*dbmodels.SecuritisationPool, error)
	SecuritisationPoolAll(ctx context.Context) ([]*dbmodels.SecuritisationPool, error)
	SecuritisationPoolUpdate(ctx context.Context, pool dbmodels.SecuritisationPool, updateCols []string) (*dbmodels.SecuritisationPool, error)
	SecuritisationPoolDelete(ctx context.Context, poolID string) error
}
