package accountlockrule

import (
	"dm_loanservice/drivers/dbmodels"
	accountlockrule "dm_loanservice/internal/service/domain/account_lock_rule"
)

func mapAccountLockRule(lockRule *dbmodels.AccountLockRule) *accountlockrule.AccountLockRule {
	return &accountlockrule.AccountLockRule{
		ID:         lockRule.ID,
		AccountID:  lockRule.AccountID,
		Status:     lockRule.AccountStatus,
		FieldName:  lockRule.FieldName,
		LockType:   lockRule.LockType,
		LockReason: lockRule.LockReason,
		CreatedAt:  lockRule.CreatedAt.String(),
		UpdatedAt:  lockRule.UpdatedAt.String(),
	}
}
