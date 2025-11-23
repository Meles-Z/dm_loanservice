package account

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
	"dm_loanservice/drivers/utils"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/drivers/uuid"
	"dm_loanservice/internal/service/domain/account"
	"dm_loanservice/internal/service/domain/dashboard"
	"fmt"
	"time"

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

func (s *svc) AccountRead(ctx context.Context, ctxDM *ctxDM.Context, req account.AccountReadRequest) (*account.AccountReadResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, utils.ErrBadRequest
	}
	acc, err := s.r.AccountRead(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	// restrict fields
	// 1. Fetch from all sources
	accountLockRules, err := s.accounLockRuleRepo.AccountLockRuleReadByAccount(ctx, acc.ID)
	if err != nil {
		return nil, err
	}

	serviceRestrictions, err := s.serviceRestrictionRepo.ServiceReadByAccount(ctx, acc.ID)
	if err != nil {
		return nil, err
	}
	isRestricted := len(accountLockRules) > 0 || len(serviceRestrictions) > 0
	restrictedFields := []string{}
	for _, r := range accountLockRules {
		field := r.FieldName
		restrictedFields = append(restrictedFields, field)
	}

	for _, r := range serviceRestrictions {
		action := r.ActionName
		restrictedFields = append(restrictedFields, action)
	}

	status := "unrestricted"
	if isRestricted {
		status = "restricted"
	}
	respMap := mapAccount(acc)
	resp := &account.AccountReadResponse{
		Account: *respMap,
		RestrictionSummary: account.RestrictionSummary{
			IsRestricted:     isRestricted,
			Status:           status,
			LockReason:       "Account is securitised",
			RestrictedFields: restrictedFields,
		},
	}

	return resp, nil
}

func (s *svc) AccountUpdate(
	ctx context.Context,
	ctxDM *ctxDM.Context,
	req *account.AccountUpdateRequest,
) (*account.AccountResponse, error) {

	// 1️⃣ Validate request
	if req.ID == "" {
		ctxDM.ErrorMessage = "account ID is required"
		return nil, utils.ErrorNewInvalidRequest
	}

	// 2️⃣ Fetch existing account
	acc, err := s.r.AccountRead(ctx, req.ID)
	if err != nil {
		ctxDM.ErrorMessage = fmt.Sprintf("account not found: %s", err.Error())
		return nil, err
	}
	// check rules
	rules, err := s.accounLockRuleRepo.AccountLockRuleRead(ctx, req.ID)
	if err != nil {
		ctxDM.ErrorMessage = err.Error()
		return nil, err
	}
	if rules.AccountStatus == "SECURITISED" {
		ctxDM.ErrorMessage = "account is securitised, cannot update"
		return nil, err
	}
	// check for the role who update the account below for later

	// 3️⃣ Apply partial updates
	updated := *acc // copy DB model
	whitelist := []string{}

	// Customer & product
	if req.MortgageID != nil {
		updated.ProductID = *req.MortgageID
		whitelist = append(whitelist, dbmodels.AccountColumns.ProductID)
	}
	if req.CustomerID != nil {
		updated.CustomerID = *req.CustomerID
		whitelist = append(whitelist, dbmodels.AccountColumns.CustomerID)
	}
	if req.ProductID != nil {
		updated.ProductID = *req.ProductID
		whitelist = append(whitelist, dbmodels.AccountColumns.ProductID)
	}

	// Financials
	if req.LoanAmount != nil {
		updated.LoanAmount = *req.LoanAmount
		whitelist = append(whitelist, dbmodels.AccountColumns.LoanAmount)
	}
	if req.BalanceOutstanding != nil {
		updated.BalanceOutstanding = *req.BalanceOutstanding
		whitelist = append(whitelist, dbmodels.AccountColumns.BalanceOutstanding)
	}

	// Dates & term
	if req.StartDate != nil {
		updated.StartDate = *req.StartDate
		whitelist = append(whitelist, dbmodels.AccountColumns.StartDate)
	}
	if req.EndDate != nil {
		updated.EndDate = *req.EndDate
		whitelist = append(whitelist, dbmodels.AccountColumns.EndDate)
	}
	if req.TermYears != nil {
		updated.TermYears = *req.TermYears
		whitelist = append(whitelist, dbmodels.AccountColumns.TermYears)
	}

	// Arrears
	if req.ArrearsFlag != nil {
		updated.ArrearsFlag = null.BoolFrom(*req.ArrearsFlag)
		whitelist = append(whitelist, dbmodels.AccountColumns.ArrearsFlag)
	}
	if req.ArrearsAmount != nil {
		updated.ArrearsAmount = types.NewNullDecimal(req.ArrearsAmount.Big)
		whitelist = append(whitelist, dbmodels.AccountColumns.ArrearsAmount)
	}
	if req.ArrearsDays != nil {
		updated.ArrearsDays = null.IntFrom(*req.ArrearsDays)
		whitelist = append(whitelist, dbmodels.AccountColumns.ArrearsDays)
	}

	// Forbearance
	if req.ForbearanceFlag != nil {
		updated.ForbearanceFlag = null.BoolFrom(*req.ForbearanceFlag)
		whitelist = append(whitelist, dbmodels.AccountColumns.ForbearanceFlag)
	}
	if req.ForbearanceType != nil {
		updated.ForbearanceType = null.StringFrom(*req.ForbearanceType)
		whitelist = append(whitelist, dbmodels.AccountColumns.ForbearanceType)
	}

	// Fraud
	if req.FraudFlag != nil {
		updated.FraudFlag = null.BoolFrom(*req.FraudFlag)
		whitelist = append(whitelist, dbmodels.AccountColumns.FraudFlag)
	}
	if req.FraudNotes != nil {
		updated.FraudNotes = null.StringFrom(*req.FraudNotes)
		whitelist = append(whitelist, dbmodels.AccountColumns.FraudNotes)
	}

	// Collateral & security
	if req.RedrawFacility != nil {
		updated.RedrawFacility = null.BoolFrom(*req.RedrawFacility)
		whitelist = append(whitelist, dbmodels.AccountColumns.RedrawFacility)
	}
	if req.CollateralAddress != nil {
		updated.CollateralAddress = null.StringFrom(*req.CollateralAddress)
		whitelist = append(whitelist, dbmodels.AccountColumns.CollateralAddress)
	}
	if req.CollateralType != nil {
		updated.CollateralType = null.StringFrom(*req.CollateralType)
		whitelist = append(whitelist, dbmodels.AccountColumns.CollateralType)
	}
	if req.SecurityType != nil {
		updated.SecurityType = null.StringFrom(*req.SecurityType)
		whitelist = append(whitelist, dbmodels.AccountColumns.SecurityType)
	}

	// Portfolio & stress
	if req.PortfolioID != nil {
		updated.PortfolioID = null.StringFrom(*req.PortfolioID)
		whitelist = append(whitelist, dbmodels.AccountColumns.PortfolioID)
	}
	if req.StressTestResult != nil {
		updated.StressTestResult = null.StringFrom(*req.StressTestResult)
		whitelist = append(whitelist, dbmodels.AccountColumns.StressTestResult)
	}
	if req.CapitalAdequacyFlag != nil {
		updated.CapitalAdequacyFlag = null.BoolFrom(*req.CapitalAdequacyFlag)
		whitelist = append(whitelist, dbmodels.AccountColumns.CapitalAdequacyFlag)
	}

	// 🚨 No fields to update
	if len(whitelist) == 0 {
		ctxDM.ErrorMessage = "no fields provided for update"
		return nil, utils.ErrorNewInvalidRequest
	}

	// 4️⃣ Persist updates
	updated.UpdatedAt = null.TimeFrom(time.Now())
	whitelist = append(whitelist, dbmodels.AccountColumns.UpdatedAt)

	_, err = s.r.AccountUpdate(ctx, updated, whitelist)
	if err != nil {
		ctxDM.ErrorMessage = fmt.Sprintf("failed to update account: %s", err.Error())
		return nil, err
	}

	// 5️⃣ Build response
	res := &account.AccountResponse{
		ID:         updated.ID,
		MortgageID: updated.ProductID,
		CustomerID: updated.CustomerID,
		ProductID:  updated.ProductID,
		StartDate:  updated.StartDate,
		EndDate:    updated.EndDate,
	}

	return res, nil
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
