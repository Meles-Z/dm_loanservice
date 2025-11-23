package investorrestriction

import (
	"context"
	"fmt"

	"dm_loanservice/drivers/dbmodels"
	domain "dm_loanservice/internal/service/domain/investor_restriction"

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

func (r *repository) InvestorRestrictionAdd(ctx context.Context, m *dbmodels.InvestorRestriction) (*dbmodels.InvestorRestriction, error) {
	err := m.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return m, nil
}
func (r *repository) InvestorRestrictionRead(ctx context.Context, investorRestrictionID string) (*dbmodels.InvestorRestriction, error) {
	investorRestriction, err := dbmodels.InvestorRestrictions(qm.Where("id = ?", investorRestrictionID)).One(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to get investor restriction by ID: %w", err)
	}
	return investorRestriction, nil
}

func (r *repository) InvestorRestrictionReadByAccount(ctx context.Context, accountID string) ([]*dbmodels.InvestorRestriction, error) {
	investorRestriction, err := dbmodels.InvestorRestrictions(qm.Where("account_id = ? AND is_active = ?", accountID, true)).All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to get investor restriction by account ID: %w", err)
	}
	return investorRestriction, nil
}
