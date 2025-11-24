package securitisation

import (
	"context"
	"fmt"

	"dm_loanservice/drivers/dbmodels"
	domain "dm_loanservice/internal/service/domain/securitisation"

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

func (r *repository) SecuritisationPoolAdd(ctx context.Context, m dbmodels.SecuritisationPool) (*dbmodels.SecuritisationPool, error) {
	err := m.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return &m, nil
}
func (r *repository) SecuritisationPoolRead(ctx context.Context, poolID string) (*dbmodels.SecuritisationPool, error) {
	pool, err := dbmodels.SecuritisationPools(qm.Where("id = ?", poolID)).One(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to get pool by ID: %w", err)
	}
	return pool, nil
}

func (r *repository) SecuritisationPoolAll(ctx context.Context) ([]*dbmodels.SecuritisationPool, error) {
	pools, err := dbmodels.SecuritisationPools(qm.OrderBy("id")).All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to get pools: %w", err)
	}
	return pools, nil
}

func (r *repository) SecuritisationPoolUpdate(ctx context.Context, pool dbmodels.SecuritisationPool, updateCols []string) (*dbmodels.SecuritisationPool, error) {

	updateCols = append(updateCols, dbmodels.SecuritisationPoolColumns.UpdatedAt)

	_, err := pool.Update(ctx, r.db, boil.Whitelist(updateCols...))
	if err != nil {
		return nil, fmt.Errorf("failed to update pool: %w", err)
	}

	return &pool, nil
}

func (r *repository) SecuritisationPoolDelete(ctx context.Context, poolID string) error {
	pool, err := dbmodels.SecuritisationPools(qm.Where("id = ?", poolID)).One(ctx, r.db)
	if err != nil {
		return fmt.Errorf("failed to get pool by ID: %w", err)
	}
	_, err = pool.Delete(ctx, r.db)
	if err != nil {
		return fmt.Errorf("failed to delete pool: %w", err)
	}
	return nil
}
