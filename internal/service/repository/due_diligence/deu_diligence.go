package duediligence

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
	domain "dm_loanservice/internal/service/domain/due_diligence"

	"github.com/aarondl/null/v8"
	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/aarondl/sqlboiler/v4/queries/qm"
	"github.com/jmoiron/sqlx"
)

func NewRepository(db *sqlx.DB) domain.Repository {
	return &repository{db: db, schema: "public"}
}

type repository struct {
	db     *sqlx.DB
	schema string
}

func (r *repository) DueDiligenceAdd(ctx context.Context, m dbmodels.Duediligencechecklistitem) (*dbmodels.Duediligencechecklistitem, error) {
	err := m.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return &m, nil
}
func (r *repository) DueDiligenceRead(ctx context.Context, id string) (*dbmodels.Duediligencechecklistitem, error) {
	duediligence, err := dbmodels.Duediligencechecklistitems(qm.Where("id = ?", id)).One(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return duediligence, nil
}

func (r *repository) DueDiligenceUpdate(
	ctx context.Context,
	req domain.DueDiligenceUpdateRequest,
	existing *dbmodels.Duediligencechecklistitem,
) (*dbmodels.Duediligencechecklistitem, error) {

	columns := []string{}

	if req.ChecklistItem != nil {
		existing.ChecklistItem = *req.ChecklistItem
		columns = append(columns, dbmodels.DuediligencechecklistitemColumns.ChecklistItem)
	}

	if req.Status != nil {
		existing.Status = *req.Status
		columns = append(columns, dbmodels.DuediligencechecklistitemColumns.Status)
	}

	if req.Comments != nil {
		existing.Comments = null.String{String: *req.Comments, Valid: true}
		columns = append(columns, dbmodels.DuediligencechecklistitemColumns.Comments)
	}

	if len(columns) == 0 {
		return existing, nil
	}

	_, err := existing.Update(ctx, r.db, boil.Whitelist(columns...))
	if err != nil {
		return nil, err
	}

	return existing, nil
}

func (r *repository) DueDiligenceByAccount(ctx context.Context, accountID string) ([]*dbmodels.Duediligencechecklistitem, error) {
	dd, err := dbmodels.Duediligencechecklistitems(qm.Where("account_id = ?", accountID)).All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return dd, nil
}
