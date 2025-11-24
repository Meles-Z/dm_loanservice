package account

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
	"dm_loanservice/internal/service/domain/dashboard"
)

type Repository interface {
	AccountAdd(ctx context.Context, m dbmodels.Account) (*dbmodels.Account, error)
	AccountRead(ctx context.Context, accountID string) (*dbmodels.Account, error)
	AccountUpdate(ctx context.Context, m dbmodels.Account, updateCols []string) (*dbmodels.Account, error)
	ListEligibleAccounts(
		ctx context.Context,
		page, pageSize int,
		mortgageType, region string,
		ltvMin, ltvMax float64,
		arrearsDaysMax int,
		originationFrom, originationTo string,
		propertyType string,
		sortBy, sortDirection string,
	) (accounts []*dbmodels.Account, total int, err error)
	AccountArrears(ctx context.Context) (*dashboard.TotalArrears, error)
	RecentArrearsCases(ctx context.Context) ([]dashboard.ArrearsCase, error)
	RecentArrears(ctx context.Context) ([]AccountRecentResponse, error)
	MortgagePerformance(ctx context.Context) (*dashboard.MortgagePerformance, error)
	AccountCount(ctx context.Context) (int64, float64, error)
}
