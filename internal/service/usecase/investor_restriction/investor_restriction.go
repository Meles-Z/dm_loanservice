package inverstoryrestriction

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/drivers/uuid"
	investorrestriction "dm_loanservice/internal/service/domain/investor_restriction"

	"github.com/aarondl/null/v8"
)

func (s *svc) InvestorRestrictionAdd(ctx context.Context, ctxDM *ctxDM.Context, req investorrestriction.InvestorRestrictionAddRequest) (*investorrestriction.InvestorRestrictionResponse, error) {
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

	invRestriction := &dbmodels.InvestorRestriction{
		ID:               uuid.UUID(),
		AccountID:        req.AccountID,
		RestrictionScope: req.RestrictionScope,
		FieldName:        null.String{String: req.FieldName, Valid: true},
		ActionName:       null.String{String: req.ActionName, Valid: true},
		RuleType:         req.RuleType,
		Reason:           req.Reason,
		IsActive:         req.IsActive,
	}
	newInvRestriction, err := s.inverstorResRepo.InvestorRestrictionAdd(ctx, invRestriction)
	if err != nil {
		return nil, err
	}
	mappedInvRestriction := mapInvestorRestriction(newInvRestriction)
	return &investorrestriction.InvestorRestrictionResponse{
		InvestorRestriction: mappedInvRestriction,
	}, nil
}

func (s *svc) InvestorRestrictionRead(ctx context.Context, ctxDM *ctxDM.Context, req investorrestriction.InvestorRestrictionReadRequest) (*investorrestriction.InvestorRestrictionResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	invRestriction, err := s.inverstorResRepo.InvestorRestrictionRead(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	mappedInvRestriction := mapInvestorRestriction(invRestriction)
	return &investorrestriction.InvestorRestrictionResponse{
		InvestorRestriction: mappedInvRestriction,
	}, nil
}
