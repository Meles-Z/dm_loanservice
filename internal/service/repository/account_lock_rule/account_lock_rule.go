package accountlockrule

import (
	"context"
	"fmt"

	"dm_loanservice/drivers/dbmodels"
	domain "dm_loanservice/internal/service/domain/account_lock_rule"

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

func (r *repository) AccountLockRuleAdd(ctx context.Context, m *dbmodels.AccountLockRule) (*dbmodels.AccountLockRule, error) {
	err := m.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return m, nil
}
func (r *repository) AccountLockRuleRead(ctx context.Context, accountLockRuleID string) (*dbmodels.AccountLockRule, error) {
	accountLockRule, err := dbmodels.AccountLockRules(qm.Where("id = ?", accountLockRuleID)).One(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to get account lock rule by ID: %w", err)
	}
	return accountLockRule, nil
}

func (r *repository) AccountLockRuleReadByAccount(ctx context.Context, accountID string) ([]*dbmodels.AccountLockRule, error) {
	accountLockRule, err := dbmodels.AccountLockRules(qm.Where("account_id = ?", accountID)).All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to get account lock rule by account ID: %w", err)
	}
	return accountLockRule, nil
}
