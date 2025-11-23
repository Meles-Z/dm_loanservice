package accountlockrule

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/drivers/uuid"
	accountlockrule "dm_loanservice/internal/service/domain/account_lock_rule"
)

func (s *svc) AccountLockRuleAdd(ctx context.Context, ctxDM *ctxDM.Context, req accountlockrule.AccountLockRuleAddRequest) (*accountlockrule.AccountLockRuleResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	acc, err := s.accountRepo.AccountRead(ctx, req.AccountID)
	if err != nil {
		ctxDM.ErrorMessage = err.Error()
		return nil, err
	}

	if acc.FraudFlag.Bool {
		return nil, err
	}

	lockRule := &dbmodels.AccountLockRule{
		ID:            uuid.UUID(),
		AccountID:     req.AccountID,
		AccountStatus: req.Status,
		FieldName:     req.FieldName,
		LockType:      req.LockType,
		LockReason:    req.LockReason,
	}
	newLockRule, err := s.accountLockRulesRepo.AccountLockRuleAdd(ctx, lockRule)
	if err != nil {
		return nil, err
	}
	mappedLockRule := mapAccountLockRule(newLockRule)
	return &accountlockrule.AccountLockRuleResponse{
		AccountLockRule: mappedLockRule,
	}, nil
}

func (s *svc) AccountLockRuleRead(ctx context.Context, ctxDM *ctxDM.Context, req accountlockrule.AccountLockRuleReadRequest) (*accountlockrule.AccountLockRuleResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	lockRule, err := s.accountLockRulesRepo.AccountLockRuleRead(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	mappedLockRule := mapAccountLockRule(lockRule)
	return &accountlockrule.AccountLockRuleResponse{
		AccountLockRule: mappedLockRule,
	}, nil
}
