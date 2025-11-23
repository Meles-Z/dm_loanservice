package servicerestriction

import (
	"context"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/internal/service/domain/account"
	accountlockrule "dm_loanservice/internal/service/domain/account_lock_rule"
	investorRestriction "dm_loanservice/internal/service/domain/investor_restriction"
	service_restriction "dm_loanservice/internal/service/domain/service_restriction"
)

type Service interface {
	ServiceRestrictionAdd(context.Context, *ctxDM.Context, service_restriction.ServiceRestrictionAddRequest) (*service_restriction.ServiceRestrictionResponse, error)
	ServiceRestrictionRead(context.Context, *ctxDM.Context, service_restriction.ServiceRestrictionReadRequest) (*service_restriction.ServiceRestrictionResponse, error)
	ServiceRestrictionReadByAccount(context.Context, *ctxDM.Context, service_restriction.ServiceRestrictionReadByAccountRequest) (*service_restriction.ServiceRestrictionSliceResponse, error)
}

func NewService(repo service_restriction.Repository,
	accountRepo account.Repository,
	investorRestriction investorRestriction.Repository,
	accountRuleRepo accountlockrule.Repository,

) Service {
	return &svc{
		serviceRestrictionRepo: repo,
		accountRepo:            accountRepo,
		investorRestriction:    investorRestriction,
		accountRuleRepo:        accountRuleRepo,
	}
}

type svc struct {
	serviceRestrictionRepo service_restriction.Repository
	accountRepo            account.Repository
	investorRestriction    investorRestriction.Repository
	accountRuleRepo        accountlockrule.Repository
}
