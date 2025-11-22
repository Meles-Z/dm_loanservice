package accountauditlog

import (
	"dm_loanservice/drivers/dbmodels"
	accountauditlog "dm_loanservice/internal/service/domain/account_audit_log"
)

func mapAccountAuditLog(log *dbmodels.AccountAuditLog) accountauditlog.AccountAuditLog {
	return accountauditlog.AccountAuditLog{
		ID:          log.ID,
		AccountID:   log.AccountID,
		Action:      log.Action.String,
		Details:     log.Details.String,
		PerformedBy: log.PerformedBy.Int,
		PerformedAt: log.PerformedAt.Time,
	}
}
