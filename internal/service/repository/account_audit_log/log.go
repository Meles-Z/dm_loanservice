package accountauditlog

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
	domain "dm_loanservice/internal/service/domain/account_audit_log"

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

func (r *repository) AccountAuditLogAdd(ctx context.Context, m dbmodels.AccountAuditLog) (*dbmodels.AccountAuditLog, error) {
	err := m.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *repository) AccountAuditLogRead(ctx context.Context, id string) (*dbmodels.AccountAuditLog, error) {
	accountAuditLog, err := dbmodels.AccountAuditLogs(qm.Where("id = ?", id)).One(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return accountAuditLog, nil
}
