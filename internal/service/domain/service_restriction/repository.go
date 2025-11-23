package servicerestriction

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
)

type Repository interface {
	ServiceRestrictionAdd(context.Context, *dbmodels.ServicingRestriction) (*dbmodels.ServicingRestriction, error)
	ServiceRestrictionRead(context.Context, string) (*dbmodels.ServicingRestriction, error)
	ServiceReadByAccount(ctx context.Context, loanID string) ([]*dbmodels.ServicingRestriction, error)
}
