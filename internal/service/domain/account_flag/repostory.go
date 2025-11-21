package accountflag

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
)

type Repository interface {
	AccountFlagAdd(ctx context.Context, m dbmodels.AccountFlag) (*dbmodels.AccountFlag, error)
	AccountFlagRead(ctx context.Context, accountID string) (*dbmodels.AccountFlag, error)
}
