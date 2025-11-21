package account

import (
	"context"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/internal/service/domain/account"
	"dm_loanservice/internal/service/domain/dashboard"
)

type Service interface {
	AccountAdd(context.Context, *ctxDM.Context, account.AccountAddRequest) (*account.AccountResponse, error)
	AccountId(context.Context, *ctxDM.Context, account.AccountReadRequest) (*account.AccountReadResponse, error)
	AccountRecentArrears(context.Context, *ctxDM.Context) ([]account.AccountRecentResponse, error)
	MortgagePerformance(context.Context, *ctxDM.Context) (*dashboard.MortgagePerformance, error)
	RecalculateArrears(ctx context.Context) error
}

func NewService(repo account.Repository) Service {
	return &svc{
		r: repo,
	}
}

type svc struct {
	r account.Repository
}
