package investorrestriction

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
)

type Repository interface {
	InvestorRestrictionAdd(context.Context, *dbmodels.InvestorRestriction) (*dbmodels.InvestorRestriction, error)
	InvestorRestrictionRead(context.Context, string) (*dbmodels.InvestorRestriction, error)
	InvestorRestrictionReadByAccount(ctx context.Context, accountID string) ([]*dbmodels.InvestorRestriction, error)
}
