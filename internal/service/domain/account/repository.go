package account

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
	"dm_loanservice/internal/service/domain/dashboard"
)

type Repository interface {
	AccountAdd(ctx context.Context, m dbmodels.Account) (*dbmodels.Account, error)
	AccountRead(ctx context.Context, accountID string) (*dbmodels.Account, error)
	AccountArrears(ctx context.Context) (*dashboard.TotalArrears, error)
	RecentArrearsCases(ctx context.Context) ([]dashboard.ArrearsCase, error)
	RecentArrears(ctx context.Context) ([]AccountRecentResponse, error)
	MortgagePerformance(ctx context.Context) (*dashboard.MortgagePerformance, error)
}
