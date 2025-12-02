package securitisation

import (
	"bytes"
	"context"
	"dm_loanservice/drivers/dbmodels"
	"dm_loanservice/drivers/utils"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/drivers/uuid"
	"dm_loanservice/internal/service/domain/securitisation"
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"strings"
	"time"

	"github.com/aarondl/null/v8"
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
		cust, err := s.customerRepo.CustomerRead(ctxDM, acc.CustomerID)
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
				FullName: cust.Customer.FirstName + " " + cust.Customer.LastName,
				Region:   cust.Customer.Address,
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
	totalLoans, _, err := s.accountRepo.AccountCount(ctx)
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

func (s *svc) SecuritisationPoolAdd(ctx context.Context, ctxDM *ctxDM.Context, req securitisation.SecuritisationPoolAddRequest) (*securitisation.SecuritisationPoolAddResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, utils.ErrBadRequest
	}

	pool := dbmodels.SecuritisationPool{
		ID:                       uuid.UUID(),
		FundingSource:            req.FundingSource,
		ServicingRole:            req.ServicingRole,
		SPVName:                  req.SPVName,
		SPVJurisdiction:          req.SPVJurisdiction,
		PoolAllocationDate:       req.PoolAllocationDate,
		LoanTransferDate:         req.LoanTransferDate,
		CurrentPoolBalance:       req.CurrentPoolBalance,
		Factor:                   req.Factor,
		NoteClass:                req.NoteClass,
		InterestRemittanceDate:   req.InterestRemittance,
		PrincipalRemittanceDate:  req.PrincipalRemittance,
		ServicingFeeRate:         req.ServicingFeeRate,
		ReportingCurrency:        req.ReportingCurrency,
		EsmaAssetCode:            null.String{String: req.ESMAAssetCode, Valid: true},
		CreditEnhancementType:    req.CreditEnhancementType,
		InvestorReportIdentifier: null.String{String: req.InvestorReportIdentifier, Valid: true},
	}

	newPool, err := s.securitisationRepo.SecuritisationPoolAdd(ctx, pool)
	if err != nil {
		return nil, err
	}
	mappedPool := mapSecuritisationPool(newPool)
	return &securitisation.SecuritisationPoolAddResponse{
		Pool: mappedPool,
	}, nil
}

func (s *svc) SecuritisationPoolRead(ctx context.Context, ctxDM *ctxDM.Context, req securitisation.SecuritisationPoolReadRequest) (*securitisation.SecuritisationPoolReadResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, utils.ErrBadRequest
	}
	pool, err := s.securitisationRepo.SecuritisationPoolRead(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	mappedPool := mapSecuritisationPool(pool)
	return &securitisation.SecuritisationPoolReadResponse{
		Pool: mappedPool,
	}, nil
}

func (s *svc) SecuritisationPoolUpdate(
	ctx context.Context,
	ctxDM *ctxDM.Context,
	req securitisation.SecuritisationPoolUpdateRequest,
) (*securitisation.SecuritisationPoolUpdateResponse, error) {

	// 1️⃣ Validate request
	if err := req.Validate(); err != nil {
		return nil, utils.ErrBadRequest
	}

	// 2️⃣ Read existing pool
	pool, err := s.securitisationRepo.SecuritisationPoolRead(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	// 3️⃣ Prepare dynamic update column list
	var updateCols []string

	// ---------- STRING FIELDS ----------
	if req.FundingSource != nil {
		pool.FundingSource = *req.FundingSource
		updateCols = append(updateCols, dbmodels.SecuritisationPoolColumns.FundingSource)
	}
	if req.ServicingRole != nil {
		pool.ServicingRole = *req.ServicingRole
		updateCols = append(updateCols, dbmodels.SecuritisationPoolColumns.ServicingRole)
	}
	if req.SPVName != nil {
		pool.SPVName = *req.SPVName
		updateCols = append(updateCols, dbmodels.SecuritisationPoolColumns.SPVName)
	}
	if req.SPVJurisdiction != nil {
		pool.SPVJurisdiction = *req.SPVJurisdiction
		updateCols = append(updateCols, dbmodels.SecuritisationPoolColumns.SPVJurisdiction)
	}
	if req.NoteClass != nil {
		pool.NoteClass = *req.NoteClass
		updateCols = append(updateCols, dbmodels.SecuritisationPoolColumns.NoteClass)
	}
	if req.ReportingCurrency != nil {
		pool.ReportingCurrency = *req.ReportingCurrency
		updateCols = append(updateCols, dbmodels.SecuritisationPoolColumns.ReportingCurrency)
	}
	if req.CreditEnhancementType != nil {
		pool.CreditEnhancementType = *req.CreditEnhancementType
		updateCols = append(updateCols, dbmodels.SecuritisationPoolColumns.CreditEnhancementType)
	}
	if req.ESMAAssetCode != nil {
		pool.EsmaAssetCode = null.String{String: *req.ESMAAssetCode, Valid: true}
		updateCols = append(updateCols, dbmodels.SecuritisationPoolColumns.EsmaAssetCode)
	}
	if req.InvestorReportIdentifier != nil {
		pool.InvestorReportIdentifier = null.String{String: *req.InvestorReportIdentifier, Valid: true}
		updateCols = append(updateCols, dbmodels.SecuritisationPoolColumns.InvestorReportIdentifier)
	}

	// ---------- DATE FIELDS ----------
	if req.PoolAllocationDate != nil {
		t, err := time.Parse(time.RFC3339, *req.PoolAllocationDate)
		if err != nil {
			return nil, utils.ErrBadRequest
		}
		pool.PoolAllocationDate = t
		updateCols = append(updateCols, dbmodels.SecuritisationPoolColumns.PoolAllocationDate)
	}

	if req.LoanTransferDate != nil {
		t, err := time.Parse(time.RFC3339, *req.LoanTransferDate)
		if err != nil {
			return nil, utils.ErrBadRequest
		}
		pool.LoanTransferDate = t
		updateCols = append(updateCols, dbmodels.SecuritisationPoolColumns.LoanTransferDate)
	}

	if req.InterestRemittance != nil {
		t, err := time.Parse(time.RFC3339, *req.InterestRemittance)
		if err != nil {
			return nil, utils.ErrBadRequest
		}
		pool.InterestRemittanceDate = t
		updateCols = append(updateCols, dbmodels.SecuritisationPoolColumns.InterestRemittanceDate)
	}

	if req.PrincipalRemittance != nil {
		t, err := time.Parse(time.RFC3339, *req.PrincipalRemittance)
		if err != nil {
			return nil, utils.ErrBadRequest
		}
		pool.PrincipalRemittanceDate = t
		updateCols = append(updateCols, dbmodels.SecuritisationPoolColumns.PrincipalRemittanceDate)
	}

	// ---------- DECIMAL FIELDS ----------
	if req.CurrentPoolBalance != nil {
		pool.CurrentPoolBalance = *req.CurrentPoolBalance
		updateCols = append(updateCols, dbmodels.SecuritisationPoolColumns.CurrentPoolBalance)
	}

	if req.Factor != nil {
		pool.Factor = *req.Factor
		updateCols = append(updateCols, dbmodels.SecuritisationPoolColumns.Factor)
	}

	if req.ServicingFeeRate != nil {
		pool.ServicingFeeRate = *req.ServicingFeeRate
		updateCols = append(updateCols, dbmodels.SecuritisationPoolColumns.ServicingFeeRate)
	}

	// 4️⃣ Save to DB
	_, err = s.securitisationRepo.SecuritisationPoolUpdate(ctx, *pool, updateCols)
	if err != nil {
		return nil, err
	}

	// 5️⃣ Map & return
	mappedPool := mapSecuritisationPool(pool)

	return &securitisation.SecuritisationPoolUpdateResponse{
		Pool: mappedPool,
	}, nil
}

func (s *svc) SecuritisationPoolAll(ctx context.Context, ctxDM *ctxDM.Context) (*securitisation.SecuritisationPoolAllResponse, error) {
	pools, err := s.securitisationRepo.SecuritisationPoolAll(ctx)
	if err != nil {
		return nil, err
	}
	mappedPools := make([]*securitisation.SecuritisationPool, len(pools))
	for i, p := range pools {
		mappedPools[i] = mapSecuritisationPool(p)
	}
	return &securitisation.SecuritisationPoolAllResponse{
		Pools: mappedPools,
	}, nil
}

func (s *svc) SecuritisationDelete(ctx context.Context, ctxDM *ctxDM.Context, req securitisation.SecuritisationDeleteRequest) (string, error) {
	if err := req.Validate(); err != nil {
		return "", utils.ErrBadRequest
	}
	return "deleted", s.securitisationRepo.SecuritisationPoolDelete(ctx, req.ID)
}

func (s *svc) SecuritisationDashboard(ctx context.Context, ctxDM *ctxDM.Context) (*securitisation.DashboardEligibleAccountResponse, error) {
	// count total loans
	totalLoans, outstanding, err := s.accountRepo.AccountCount(ctx)
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

	return &securitisation.DashboardEligibleAccountResponse{
		TotalLoans:       totalLoans,
		EligibleLoans:    eligible,
		IneligibleLoans:  ineligible,
		TotalOutstanding: outstanding,
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

func (s *svc) SecuritisationDashboardExport(
	ctx context.Context,
	ctxDM *ctxDM.Context,
	req securitisation.DashboardExportRequest,
) (*securitisation.DashboardExportResponse, error) {

	// 1️⃣ Validate format
	format := strings.ToLower(req.Format)
	if format != "csv" && format != "xlsx" {
		return nil, fmt.Errorf("invalid format: must be csv or xlsx")
	}

	// 2️⃣ Load dashboard data using your existing function
	dashboard, err := s.SecuritisationDashboard(ctx, ctxDM)
	if err != nil {
		return nil, err
	}

	// 3️⃣ Generate export based on format
	switch format {

	case "csv":
		fileBytes, err := generateDashboardCSV(dashboard)
		if err != nil {
			ctxDM.ErrorMessage = fmt.Sprintf("csv generation failed: %s", err.Error())
			return nil, err
		}

		return &securitisation.DashboardExportResponse{
			FileName: "securitisation_dashboard.csv",
			MimeType: "text/csv",
			FileData: base64.StdEncoding.EncodeToString(fileBytes),
		}, nil

	case "xlsx":
		fileBytes, err := generateDashboardXLSX(dashboard)
		if err != nil {
			ctxDM.ErrorMessage = fmt.Sprintf("xlsx generation failed: %s", err.Error())
			return nil, err
		}

		return &securitisation.DashboardExportResponse{
			FileName: "securitisation_dashboard.xlsx",
			MimeType: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
			FileData: base64.StdEncoding.EncodeToString(fileBytes),
		}, nil
	}

	return nil, nil
}

func (s *svc) SecuritisationPoolReport(ctx context.Context, ctxDM *ctxDM.Context, req securitisation.SecuritisationPoolReportRequest) (*securitisation.SecuritisationPoolReportResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, utils.ErrBadRequest
	}
	pool, err := s.securitisationRepo.SecuritisationPoolRead(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	// count total loans
	totalLoans, outstanding, err := s.accountRepo.AccountCount(ctx)
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

	return &securitisation.SecuritisationPoolReportResponse{
		PoolID:             pool.ID,
		SPVName:            pool.SPVName,
		FundingSource:      pool.FundingSource,
		PoolAllocationDate: pool.PoolAllocationDate.String(),
		LoanTransferDate:   pool.LoanTransferDate.String(),
		NoteClass:          pool.NoteClass,
		ReportingCurrency:  pool.ReportingCurrency,
		//
		TotalLoans:       totalLoans,
		EligibleLoans:    eligible,
		IneligibleLoans:  ineligible,
		TotalOutstanding: outstanding,
		//
		DDSummary: securitisation.DueDiligenceSummary{
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
