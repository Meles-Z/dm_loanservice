package securitisation

import (
	"context"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/internal/service/domain/account"
	accountflag "dm_loanservice/internal/service/domain/account_flag"
	"dm_loanservice/internal/service/domain/collateral"
	"dm_loanservice/internal/service/domain/customer"
	duediligence "dm_loanservice/internal/service/domain/due_diligence"
	"dm_loanservice/internal/service/domain/securitisation"
)

type Service interface {
	EligibleAccount(ctx context.Context, ctxDM *ctxDM.Context, req securitisation.EligibleLoansQuery) (*securitisation.EligibleAccountResponse, error)
	EligibleSummary(ctx context.Context, ctxDM *ctxDM.Context) (*securitisation.EligibleAccountSummaryResponse, error)
	EligibleAccountSummaryReport(ctx context.Context, ctxDM *ctxDM.Context, req securitisation.EligibleLoansQuery) (*securitisation.EligibleAccountSummaryReportResponse, error)
}

func NewService(
	accountRepo account.Repository,
	customerRepo customer.Repository,
	dueDiligenceRepo duediligence.Repository,
	accountFlagRepo accountflag.Repository,
	collateralRepo collateral.Repository,
) Service {
	return &svc{
		accountRepo:      accountRepo,
		customerRepo:     customerRepo,
		dueDiligenceRepo: dueDiligenceRepo,
		accountflagRepo:  accountFlagRepo,
		collateralRepo:   collateralRepo,
	}
}

type svc struct {
	accountRepo      account.Repository
	customerRepo     customer.Repository
	dueDiligenceRepo duediligence.Repository
	accountflagRepo  accountflag.Repository
	collateralRepo   collateral.Repository
}
