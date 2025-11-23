package servicerestriction

import (
	"context"
	"fmt"

	"dm_loanservice/drivers/dbmodels"
	domain "dm_loanservice/internal/service/domain/service_restriction"

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

func (r *repository) ServiceRestrictionAdd(ctx context.Context, m *dbmodels.ServicingRestriction) (*dbmodels.ServicingRestriction, error) {
	err := m.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return m, nil
}
func (r *repository) ServiceRestrictionRead(ctx context.Context, serviceRestrictionID string) (*dbmodels.ServicingRestriction, error) {
	serviceRestriction, err := dbmodels.ServicingRestrictions(qm.Where("id = ?", serviceRestrictionID)).One(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to get service restriction by ID: %w", err)
	}
	return serviceRestriction, nil
}

func (r *repository) ServiceReadByAccount(ctx context.Context, accountID string) ([]*dbmodels.ServicingRestriction, error) {
	serviceRestriction, err := dbmodels.ServicingRestrictions(qm.Where("account_id = ? AND is_active = ?", accountID, true)).All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to get service restriction by account ID: %w", err)
	}
	return serviceRestriction, nil
}
