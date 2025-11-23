package duediligence

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
)

type Repository interface {
	DueDiligenceAdd(context.Context, dbmodels.Duediligencechecklistitem) (*dbmodels.Duediligencechecklistitem, error)
	DueDiligenceRead(context.Context, string) (*dbmodels.Duediligencechecklistitem, error)
	DueDiligenceUpdate(context.Context, DueDiligenceUpdateRequest, *dbmodels.Duediligencechecklistitem) (*dbmodels.Duediligencechecklistitem, error)
	DueDiligenceByAccount(context.Context, string) ([]*dbmodels.Duediligencechecklistitem, error)
	DueDiligenceStatusSummary(context.Context) (int64, int64, int64, error)
	GetDDStatusMap(ctx context.Context) (map[string]string, error)
}
