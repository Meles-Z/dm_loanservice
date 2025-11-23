package collateral

import (
	"context"
	"fmt"

	"dm_loanservice/drivers/dbmodels"
	domain "dm_loanservice/internal/service/domain/collateral"

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

func (r *repository) CollateralAdd(ctx context.Context, m dbmodels.Collateral) (*dbmodels.Collateral, error) {
	err := m.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return &m, nil
}
func (r *repository) CollateralRead(ctx context.Context, collateralID string) (*dbmodels.Collateral, error) {
	collateral, err := dbmodels.Collaterals(qm.Where("id = ?", collateralID)).One(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to get collateral by ID: %w", err)
	}
	return collateral, nil
}
