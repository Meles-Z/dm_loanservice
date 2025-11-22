package accountauditlog

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
)

type Repository interface {
	AccountAuditLogAdd(ctx context.Context, m dbmodels.AccountAuditLog) (*dbmodels.AccountAuditLog, error)
	AccountAuditLogRead(ctx context.Context, id string) (*dbmodels.AccountAuditLog, error)
}
