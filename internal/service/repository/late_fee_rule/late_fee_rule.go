package latefeerule

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
	domain "dm_loanservice/internal/service/domain/late_fee_rule"
	"fmt"

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

func (r *repository) LetFeeRuleAdd(ctx context.Context, rule dbmodels.LateFeeRule) (*dbmodels.LateFeeRule, error) {
	err := rule.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return &rule, nil
}

func (r *repository) LetFeeRuleRead(ctx context.Context, id string) (*dbmodels.LateFeeRule, error) {
	rule, err := dbmodels.LateFeeRules(qm.Where("id=?", id)).One(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return rule, nil
}

func (r *repository) LetFeeRuleReadByProduct(ctx context.Context, productID string) ([]*dbmodels.LateFeeRule, error) {
	rules, err := dbmodels.LateFeeRules(
		qm.Where("product_id = ?", productID),
	).All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch late fee rules: %w", err)
	}
	return rules, nil
}

func (r *repository) LateFeeRuleUpdate(ctx context.Context, rule dbmodels.LateFeeRule) (*dbmodels.LateFeeRule, error) {
	_, err := rule.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return &rule, nil
}
