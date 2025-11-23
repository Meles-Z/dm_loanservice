package accountflag

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
)

type Repository interface {
	AccountFlagAdd(ctx context.Context, m dbmodels.AccountFlag) (*dbmodels.AccountFlag, error)
	AccountFlagRead(ctx context.Context, accountID string) (*dbmodels.AccountFlag, error)
	AccountFlagReadByAccountId(ctx context.Context, accountID string) ([]*dbmodels.AccountFlag, error)
	AccountFlagSummary(ctx context.Context) (int64, int64, int64, error)
	GetFlagStatusMap(ctx context.Context) (map[string]string, error)
}
