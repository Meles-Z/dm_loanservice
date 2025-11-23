package inverstoryrestriction

import (
	"dm_loanservice/drivers/dbmodels"
	investorrestriction "dm_loanservice/internal/service/domain/investor_restriction"
)

func mapInvestorRestriction(invRestriction *dbmodels.InvestorRestriction) *investorrestriction.InvestorRestriction {
	return &investorrestriction.InvestorRestriction{
		ID:               invRestriction.ID,
		AccountID:        invRestriction.AccountID,
		RestrictionScope: invRestriction.RestrictionScope,
		FieldName:        invRestriction.FieldName.String,
		ActionName:       invRestriction.ActionName.String,
		RuleType:         invRestriction.RuleType,
		Reason:           invRestriction.Reason,
		IsActive:         invRestriction.IsActive,
		CreatedAt:        invRestriction.CreatedAt.String(),
		UpdatedAt:        invRestriction.UpdatedAt.String(),
	}
}
