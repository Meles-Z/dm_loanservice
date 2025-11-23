package securitisation

import (
	"bytes"
	"context"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/internal/service/domain/securitisation"
	"encoding/base64"
	"encoding/csv"
	"fmt"
)

func (s *svc) EligibleAccount(
	ctx context.Context,
	ctxDM *ctxDM.Context,
	req securitisation.EligibleLoansQuery,
) (*securitisation.EligibleAccountResponse, error) {

	// 1️⃣ Fetch filtered + paginated accounts
	accounts, total, err := s.accountRepo.ListEligibleAccounts(
		ctx,
		req.Page,
		req.PageSize,
		req.MortgageType,
		req.Region,
		req.LTVMin,
		req.LTVMax,
		req.ArrearsDaysMax,
		req.OriginationFrom,
		req.OriginationTo,
		req.PropertyType,
		req.SortBy,
		req.SortDirection,
	)
	if err != nil {
		ctxDM.ErrorMessage = "failed to list eligible accounts"
		return nil, err
	}

	var loans []securitisation.EligibleLoanItem

	// 2️⃣ Process each account
	for _, acc := range accounts {

		// 2.1 Fetch Borrower
		cust, err := s.customerRepo.FindCustomerById(ctx, acc.CustomerID)
		if err != nil {
			ctxDM.Logger.Warn(ctx, "customer not found for account")
			continue
		}

		// 2.2 Fetch Collateral / Property
		col, err := s.collateralRepo.CollateralRead(ctx, acc.ID)
		if err != nil {
			ctxDM.Logger.Warn(ctx, "collateral not found for account %s")
			continue
		}

		// 2.3 Fetch Due Diligence
		dd, err := s.dueDiligenceRepo.DueDiligenceRead(ctx, acc.ID)
		if err != nil {
			ctxDM.Logger.Warn(ctx, "due diligence not found for account %s")
			continue
		}

		// 2.4 Fetch Flags
		flags, err := s.accountflagRepo.AccountFlagReadByAccountId(ctx, acc.ID)
		if err != nil {
			ctxDM.Logger.Warn(ctx, "flags not found for account %s")
			continue
		}

		var flagNames []string
		for _, f := range flags {
			flagNames = append(flagNames, f.FlagType)
		}

		// 2.5 Compute Securitisation Eligibility
		eligibility := computeEligibility(dd.Status, flagNames)

		// 2.6 Build final loan item
		item := securitisation.EligibleLoanItem{
			LoanID: acc.ID,
			// AccountNumber: acc,
			Borrower: securitisation.BorrowerInfo{
				FullName: cust.FirstName + " " + cust.LastName,
				Region:   cust.Address.String,
			},
			Property: securitisation.PropertyInfo{
				PropertyType: col.PropertyID,
				// Region:       col.,
				// Valuation:    col.Valuation,
			},
			// LTV:             acc.LTV,
			// ArrearsDays:     acc.ArrearsDays,
			// OriginationDate: acc.OriginationDate,
			DDStatus:          dd.Status,
			EligibilityStatus: eligibility,
			Flags:             flagNames,
		}

		loans = append(loans, item)
	}

	// 3️⃣ Return Response
	return &securitisation.EligibleAccountResponse{
		Page:         req.Page,
		PageSize:     req.PageSize,
		TotalRecords: total,
		Loans:        loans,
	}, nil
}

func (s *svc) EligibleSummary(ctx context.Context, ctxDM *ctxDM.Context) (*securitisation.EligibleAccountSummaryResponse, error) {

	// count total loans
	totalLoans, err := s.accountRepo.AccountCount(ctx)
	if err != nil {
		ctxDM.ErrorMessage = fmt.Sprintf("failed to count accounts: %s", err.Error())
		return nil, err
	}

	// DD summary
	ddPass, ddPending, ddFail, err := s.dueDiligenceRepo.DueDiligenceStatusSummary(ctx)
	if err != nil {
		ctxDM.ErrorMessage = "failed to get DD summary"
		return nil, err
	}

	// flag summary
	secReady, secExcluded, manualReview, err := s.accountflagRepo.AccountFlagSummary(ctx)
	if err != nil {
		ctxDM.ErrorMessage = "failed to get flag summary"
		return nil, err
	}

	// eligibility logic (business-approved)
	eligible := ddPass
	eligible -= ddFail + ddPending
	eligible -= secExcluded + manualReview
	if eligible < 0 {
		eligible = 0
	}

	ineligible := totalLoans - eligible

	return &securitisation.EligibleAccountSummaryResponse{
		TotalLoans:      int(totalLoans),
		EligibleLoans:   int(eligible),
		IneligibleLoans: int(ineligible),
		DueDiligenceSummary: securitisation.DueDiligenceSummary{
			Pass:    int(ddPass),
			Fail:    int(ddFail),
			Pending: int(ddPending),
		},
		FlagSummary: securitisation.FlagSummary{
			SecReady:     int(secReady),
			SecExcluded:  int(secExcluded),
			ManualReview: int(manualReview),
		},
	}, nil
}

func (s *svc) EligibleAccountSummaryReport(
	ctx context.Context,
	ctxDM *ctxDM.Context,
	req securitisation.EligibleLoansQuery,
) (*securitisation.EligibleAccountSummaryReportResponse, error) {

	// 1️⃣ Fetch accounts with filters applied
	loans, _, err := s.accountRepo.ListEligibleAccounts(ctx,
		req.Page,
		req.PageSize,
		req.MortgageType,
		req.Region,
		req.LTVMin,
		req.LTVMax,
		req.ArrearsDaysMax,
		req.OriginationFrom,
		req.OriginationTo,
		req.PropertyType,
		req.SortBy,
		req.SortDirection)
	if err != nil {
		ctxDM.ErrorMessage = fmt.Sprintf("failed to fetch filtered loans: %s", err.Error())
		return nil, err
	}

	if len(loans) == 0 {
		return &securitisation.EligibleAccountSummaryReportResponse{
			File: "",
		}, nil
	}

	// 2️⃣ Fetch DD statuses for all accounts
	ddMap, err := s.dueDiligenceRepo.GetDDStatusMap(ctx)
	if err != nil {
		ctxDM.ErrorMessage = "failed to load due diligence statuses"
		return nil, err
	}

	// 3️⃣ Fetch flags for all accounts
	flagMap, err := s.accountflagRepo.GetFlagStatusMap(ctx)
	if err != nil {
		ctxDM.ErrorMessage = "failed to load account flags"
		return nil, err
	}

	// 4️⃣ Create CSV
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	// CSV Header
	writer.Write([]string{
		"AccountID",
		"AccountNumber",
		"CustomerName",
		"Product",
		"OutstandingBalance",
		"DDStatus",
		"FlagStatus",
		"EligibilityStatus",
	})

	// 5️⃣ Loop each loan and determine final eligibility
	for _, loan := range loans {

		ddStatus := ddMap[loan.ID]     // pass/pending/fail
		flagStatus := flagMap[loan.ID] // sec_ready / manual_review / sec_excluded
		eligibility := calculateEligibility(ddStatus, flagStatus)

		writer.Write([]string{
			loan.ID,
			"",
			loan.GetCustomer().FirstName + " " + loan.GetCustomer().LastName,
			loan.GetProduct().ProductName,
			loan.BalanceOutstanding.String(),
			ddStatus,
			flagStatus,
			eligibility,
		})
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		ctxDM.ErrorMessage = fmt.Sprintf("csv writing error: %s", err.Error())
		return nil, err
	}

	// 6️⃣ Base64 encode CSV file
	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())

	return &securitisation.EligibleAccountSummaryReportResponse{
		File: encoded,
	}, nil
}
