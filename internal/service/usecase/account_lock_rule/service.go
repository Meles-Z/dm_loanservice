package accountlockrule

import (
	"context"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/internal/service/domain/account"
	accountlockrule "dm_loanservice/internal/service/domain/account_lock_rule"
)

type Service interface {
	AccountLockRuleAdd(context.Context, *ctxDM.Context, accountlockrule.AccountLockRuleAddRequest) (*accountlockrule.AccountLockRuleResponse, error)
	AccountLockRuleRead(context.Context, *ctxDM.Context, accountlockrule.AccountLockRuleReadRequest) (*accountlockrule.AccountLockRuleResponse, error)
}

func NewService(repo accountlockrule.Repository, accountRepo account.Repository) Service {
	return &svc{
		accountLockRulesRepo: repo,
		accountRepo:          accountRepo,
	}
}

type svc struct {
	accountLockRulesRepo accountlockrule.Repository
	accountRepo          account.Repository
}
