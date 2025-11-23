package property

import (
	"context"
	"fmt"

	"dm_loanservice/drivers/dbmodels"
	domain "dm_loanservice/internal/service/domain/property"

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

func (r *repository) PropertyAdd(ctx context.Context, m dbmodels.Property) (*dbmodels.Property, error) {
	err := m.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return &m, nil
}
func (r *repository) PropertyRead(ctx context.Context, propertyID string) (*dbmodels.Property, error) {
	property, err := dbmodels.Properties(qm.Where("id = ?", propertyID)).One(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to get property by ID: %w", err)
	}
	return property, nil
}
