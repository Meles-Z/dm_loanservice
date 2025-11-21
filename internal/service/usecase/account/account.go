package account

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/drivers/uuid"
	"dm_loanservice/internal/service/domain/account"
	"dm_loanservice/internal/service/domain/dashboard"

	"github.com/aarondl/null/v8"
	"github.com/aarondl/sqlboiler/v4/types"
)

func (s *svc) AccountAdd(ctx context.Context, ctxdm *ctxDM.Context, req account.AccountAddRequest) (*account.AccountResponse, error) {
	// Map request → DB model
	acc := dbmodels.Account{
		ID:                 uuid.UUID(),
		CustomerID:         req.CustomerID,
		ProductID:          req.ProductID,
		LoanAmount:         req.LoanAmount,
		BalanceOutstanding: req.BalanceOutstanding,
		StartDate:          req.StartDate,
		EndDate:            req.EndDate,
		TermYears:          req.TermYears,

		ArrearsFlag:   null.BoolFrom(req.ArrearsFlag),
		ArrearsAmount: types.NullDecimal(req.ArrearsAmount),
		ArrearsDays:   null.IntFrom(req.ArrearsDays),

		ForbearanceFlag: null.BoolFrom(req.ForbearanceFlag),
		ForbearanceType: null.StringFrom(string(req.ForbearanceType)),

		FraudFlag:  null.BoolFrom(req.FraudFlag),
		FraudNotes: null.StringFrom(req.FraudNotes),

		RedrawFacility:    null.BoolFrom(req.RedrawFacility),
		CollateralAddress: null.StringFrom(req.CollateralAddress),
		CollateralType:    null.StringFrom(string(req.CollateralType)),
		SecurityType:      null.StringFrom(string(req.SecurityType)),

		PortfolioID:         null.StringFrom(req.PortfolioID),
		StressTestResult:    null.StringFrom(string(req.StressTestResult)),
		CapitalAdequacyFlag: null.BoolFrom(req.CapitalAdequacyFlag),
	}

	// Save to repository
	c, err := s.r.AccountAdd(ctx, acc)
	if err != nil {
		return nil, err
	}

	// Map DB model → Response DTO
	resp := &account.AccountResponse{
		ID:         c.ID,
		CustomerID: c.CustomerID,
		ProductID:  c.ProductID,
		StartDate:  c.StartDate,
		EndDate:    c.EndDate,
	}

	return resp, nil
}

func (s *svc) AccountId(ctx context.Context, ctxDM *ctxDM.Context, req account.AccountReadRequest) (*account.AccountReadResponse, error) {
	acc, err := s.r.AccountId(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	respMap := mapAccount(acc)
	resp := &account.AccountReadResponse{
		Account: *respMap,
	}

	return resp, nil
}
func (s *svc) AccountRecentArrears(ctx context.Context, ctcDM *ctxDM.Context) ([]account.AccountRecentResponse, error) {
	arr, err := s.r.RecentArrears(ctx)
	if err != nil {
		return nil, err
	}
	return arr, nil
}
func (s *svc) MortgagePerformance(ctx context.Context, ctxDM *ctxDM.Context) (*dashboard.MortgagePerformance, error) {
	summary, err := s.r.MortgagePerformance(ctx)
	if err != nil {
		return nil, err
	}
	return summary, nil
}

func (s *svc) RecalculateArrears(ctx context.Context) error {
	_, err := s.r.AccountArrears(ctx)
	return err
}
