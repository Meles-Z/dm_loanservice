package account

import (
	"context"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/internal/service/domain/account"
	accountlockrule "dm_loanservice/internal/service/domain/account_lock_rule"
	"dm_loanservice/internal/service/domain/dashboard"
	duediligence "dm_loanservice/internal/service/domain/due_diligence"
	servicerestriction "dm_loanservice/internal/service/domain/service_restriction"
)

type Service interface {
	AccountAdd(context.Context, *ctxDM.Context, account.AccountAddRequest) (*account.AccountResponse, error)
	AccountRead(context.Context, *ctxDM.Context, account.AccountReadRequest) (*account.AccountReadResponse, error)
	AccountUpdate(context.Context, *ctxDM.Context, *account.AccountUpdateRequest) (*account.AccountResponse, error)
	AccountRecentArrears(context.Context, *ctxDM.Context) ([]account.AccountRecentResponse, error)
	MortgagePerformance(context.Context, *ctxDM.Context) (*dashboard.MortgagePerformance, error)
	RecalculateArrears(ctx context.Context) error
}

func NewService(repo account.Repository, duediligenceRepo duediligence.Repository,
	accountLockRule accountlockrule.Repository,
	serviceRestrictionRepo servicerestriction.Repository,
) Service {
	return &svc{
		r:                      repo,
		duediligenceRepo:       duediligenceRepo,
		accounLockRuleRepo:     accountLockRule,
		serviceRestrictionRepo: serviceRestrictionRepo,
	}
}

type svc struct {
	r                      account.Repository
	duediligenceRepo       duediligence.Repository
	accounLockRuleRepo     accountlockrule.Repository
	serviceRestrictionRepo servicerestriction.Repository
}
