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
