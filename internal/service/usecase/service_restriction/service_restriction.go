package servicerestriction

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/drivers/uuid"
	domain "dm_loanservice/internal/service/domain/service_restriction"

	"github.com/aarondl/null/v8"
)

func (s *svc) ServiceRestrictionAdd(ctx context.Context, ctxDM *ctxDM.Context, req domain.ServiceRestrictionAddRequest) (*domain.ServiceRestrictionResponse, error) {
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

	serviceRest := &dbmodels.ServicingRestriction{
		ID:              uuid.UUID(),
		AccountID:       null.String{String: req.AccountID, Valid: true},
		RestrictionType: req.RestrictionType,
		ActionName:      req.ActionName,
		IsActive:        req.IsActive,
		Reason:          req.Reason,
		Source:          null.String{String: req.Source, Valid: true},
	}
	newServiceRestriction, err := s.serviceRestrictionRepo.ServiceRestrictionAdd(ctx, serviceRest)
	if err != nil {
		return nil, err
	}
	mappedServiceRestriction := mapServiceRestriction(newServiceRestriction)
	return &domain.ServiceRestrictionResponse{
		ServiceRestriction: mappedServiceRestriction,
	}, nil
}

func (s *svc) ServiceRestrictionRead(ctx context.Context, ctxDM *ctxDM.Context, req domain.ServiceRestrictionReadRequest) (*domain.ServiceRestrictionResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	serviceRest, err := s.serviceRestrictionRepo.ServiceRestrictionRead(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	mappedServiceRestriction := mapServiceRestriction(serviceRest)
	return &domain.ServiceRestrictionResponse{
		ServiceRestriction: mappedServiceRestriction,
	}, nil
}

func (s *svc) ServiceRestrictionReadByAccount(ctx context.Context, ctxDM *ctxDM.Context, req domain.ServiceRestrictionReadByAccountRequest) (*domain.ServiceRestrictionSliceResponse, error) {

	if err := req.Validate(); err != nil {
		return nil, err
	}
	// 1. Fetch from all sources
	accountLockRules, err := s.accountRuleRepo.AccountLockRuleReadByAccount(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	serviceRestrictions, err := s.serviceRestrictionRepo.ServiceReadByAccount(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	investorRestrictions, err := s.investorRestriction.InvestorRestrictionReadByAccount(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	// 2. Create unified list
	restrictions := make([]domain.RestrictionItem, 0)

	// ---- Loan Lock Rules ----
	for _, r := range accountLockRules {
		f := r.FieldName
		restrictions = append(restrictions, domain.RestrictionItem{
			ID:               r.ID,
			Type:             "loan_lock_rule",
			RestrictionLevel: r.LockType,
			FieldName:        &f,
			Reason:           r.LockReason,
			Source:           "system",
		})
	}

	// ---- Servicing Restrictions ----
	for _, r := range serviceRestrictions {
		a := r.ActionName
		restrictions = append(restrictions, domain.RestrictionItem{
			ID:               r.ID,
			Type:             "servicing_restriction",
			RestrictionLevel: "soft", // default for servicing
			ActionName:       &a,
			Reason:           r.Reason,
			Source:           r.Source.String,
		})
	}

	// ---- Investor Restrictions ----
	for _, r := range investorRestrictions {
		var field *string
		var action *string

		if r.FieldName.Valid {
			field = &r.FieldName.String
		}
		if r.ActionName.Valid {
			action = &r.ActionName.String
		}

		inv := r.AccountID
		restrictions = append(restrictions, domain.RestrictionItem{
			ID:               r.ID,
			Type:             "investor_restriction",
			RestrictionLevel: r.RuleType,
			FieldName:        field,
			ActionName:       action,
			Reason:           r.Reason,
			Source:           "investor", // TODO: investor source
			InvestorID:       &inv,
		})
	}

	// 3. Calculate status
	status := "unrestricted"
	if len(restrictions) > 0 {
		status = "restricted"
	}

	// 4. Final aggregated response
	return &domain.ServiceRestrictionSliceResponse{
		OverallRestrictionStatus: status,
		Restrictions:             restrictions,
	}, nil
}
