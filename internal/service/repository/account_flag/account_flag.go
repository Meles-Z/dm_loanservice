package accountflag

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
	domain "dm_loanservice/internal/service/domain/account_flag"

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

func (r *repository) AccountFlagAdd(ctx context.Context, m dbmodels.AccountFlag) (*dbmodels.AccountFlag, error) {
	err := m.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return &m, nil
}
func (r *repository) AccountFlagRead(ctx context.Context, id string) (*dbmodels.AccountFlag, error) {
	accountFlag, err := dbmodels.AccountFlags(qm.Where("id = ?", id)).One(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return accountFlag, nil
}

func (r *repository) AccountFlagReadByAccountId(ctx context.Context, accountID string) ([]*dbmodels.AccountFlag, error) {
	accountFlags, err := dbmodels.AccountFlags(qm.Where("account_id = ?", accountID)).All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return accountFlags, nil
}

func (r *repository) AccountFlagSummary(ctx context.Context) (secReady, secExcluded, manualReview int64, err error) {
	query := `
        SELECT flag_type, COUNT(DISTINCT account_id)
        FROM account_flags
        GROUP BY flag_type;
    `
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return
	}
	defer rows.Close()

	var flagType string
	var count int64

	for rows.Next() {
		if err = rows.Scan(&flagType, &count); err != nil {
			return
		}
		switch flagType {
		case "sec_ready":
			secReady = count
		case "sec_excluded":
			secExcluded = count
		case "manual_review":
			manualReview = count
		}
	}
	return
}

func (r *repository) GetFlagStatusMap(ctx context.Context) (map[string]string, error) {

	query := `
        SELECT account_id, flag_type
        FROM account_flags
        WHERE deleted_at IS NULL
        QUALIFY ROW_NUMBER() OVER (PARTITION BY account_id ORDER BY created_at DESC) = 1;
    `

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]string)

	var accountID, flag string

	for rows.Next() {
		if err := rows.Scan(&accountID, &flag); err != nil {
			return nil, err
		}
		result[accountID] = flag
	}

	return result, nil
}
