package inverstoryrestriction

import (
	"context"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/internal/service/domain/account"
	investorrestriction "dm_loanservice/internal/service/domain/investor_restriction"
)

type Service interface {
	InvestorRestrictionAdd(context.Context, *ctxDM.Context, investorrestriction.InvestorRestrictionAddRequest) (*investorrestriction.InvestorRestrictionResponse, error)
	InvestorRestrictionRead(context.Context, *ctxDM.Context, investorrestriction.InvestorRestrictionReadRequest) (*investorrestriction.InvestorRestrictionResponse, error)
}

func NewService(repo investorrestriction.Repository, accountRepo account.Repository) Service {
	return &svc{
		inverstorResRepo: repo,
		accountRepo:      accountRepo,
	}
}

type svc struct {
	inverstorResRepo investorrestriction.Repository
	accountRepo      account.Repository
}
