package account

import (
	"context"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/internal/service/domain/account"
	"dm_loanservice/internal/service/domain/dashboard"
	duediligence "dm_loanservice/internal/service/domain/due_diligence"
)

type Service interface {
	AccountAdd(context.Context, *ctxDM.Context, account.AccountAddRequest) (*account.AccountResponse, error)
	AccountRead(context.Context, *ctxDM.Context, account.AccountReadRequest) (*account.AccountReadResponse, error)
	AccountRecentArrears(context.Context, *ctxDM.Context) ([]account.AccountRecentResponse, error)
	MortgagePerformance(context.Context, *ctxDM.Context) (*dashboard.MortgagePerformance, error)
	RecalculateArrears(ctx context.Context) error
}

func NewService(repo account.Repository, duediligenceRepo duediligence.Repository) Service {
	return &svc{
		r:                repo,
		duediligenceRepo: duediligenceRepo,
	}
}

type svc struct {
	r                account.Repository
	duediligenceRepo duediligence.Repository
}
