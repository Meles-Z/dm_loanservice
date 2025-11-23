package accountlockrule

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
)

type Repository interface {
	AccountLockRuleAdd(context.Context, *dbmodels.AccountLockRule) (*dbmodels.AccountLockRule, error)
	AccountLockRuleRead(context.Context, string) (*dbmodels.AccountLockRule, error)
	AccountLockRuleReadByAccount(ctx context.Context, accountID string) ([]*dbmodels.AccountLockRule, error)
}
