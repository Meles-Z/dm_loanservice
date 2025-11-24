package securitisation

import (
	"bytes"
	"dm_loanservice/drivers/dbmodels"
	"dm_loanservice/internal/service/domain/securitisation"
	"encoding/csv"
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func computeEligibility(ddStatus string, flags []string) string {
	// ❌ If DD failed → not eligible
	if ddStatus == "fail" {
		return "not_eligible"
	}

	// ❌ If any blocking flag exists → not eligible
	for _, f := range flags {
		if f == "ineligible" {
			return "not_eligible"
		}
	}

	// ❌ If DD pending → review
	if ddStatus == "pending" {
		return "under_review"
	}

	// ✅ Passed everything
	return "eligible"
}

func calculateEligibility(ddStatus, flagStatus string) string {
	if ddStatus != "pass" {
		return "ineligible"
	}
	if flagStatus == "sec_excluded" || flagStatus == "manual_review" {
		return "ineligible"
	}
	return "eligible"
}

func mapSecuritisationPool(pool *dbmodels.SecuritisationPool) *securitisation.SecuritisationPool {
	return &securitisation.SecuritisationPool{
		ID:                    pool.ID,
		FundingSource:         pool.FundingSource,
		ServicingRole:         pool.ServicingRole,
		SPVName:               pool.SPVName,
		SPVJurisdiction:       pool.SPVJurisdiction,
		PoolAllocationDate:    pool.PoolAllocationDate.String(),
		LoanTransferDate:      pool.LoanTransferDate.String(),
		CurrentPoolBalance:    pool.CurrentPoolBalance,
		Factor:                pool.Factor,
		NoteClass:             pool.NoteClass,
		InterestRemittance:    pool.InterestRemittanceDate.String(),
		PrincipalRemittance:   pool.PrincipalRemittanceDate.String(),
		ServicingFeeRate:      pool.ServicingFeeRate,
		ESMAAssetCode:         pool.EsmaAssetCode.String,
		ReportingCurrency:     pool.ReportingCurrency,
		CreditEnhancementType: pool.CreditEnhancementType,
		CreatedAt:             pool.CreatedAt.String(),
		UpdatedAt:             pool.UpdatedAt.String(),
	}
}

func generateDashboardCSV(d *securitisation.DashboardEligibleAccountResponse) ([]byte, error) {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	// Header
	writer.Write([]string{"Metric", "Value"})

	// Body
	writer.Write([]string{"Total Loans", strconv.FormatInt(d.TotalLoans, 10)})
	writer.Write([]string{"Eligible Loans", strconv.FormatInt(d.EligibleLoans, 10)})
	writer.Write([]string{"Ineligible Loans", strconv.FormatInt(d.IneligibleLoans, 10)})
	writer.Write([]string{"Total Outstanding", strconv.FormatFloat(d.TotalOutstanding, 'f', 2, 64)})

	writer.Write([]string{"DD Pass", strconv.Itoa(d.DueDiligenceSummary.Pass)})
	writer.Write([]string{"DD Fail", strconv.Itoa(d.DueDiligenceSummary.Fail)})
	writer.Write([]string{"DD Pending", strconv.Itoa(d.DueDiligenceSummary.Pending)})

	writer.Write([]string{"SEC Ready", strconv.Itoa(d.FlagSummary.SecReady)})
	writer.Write([]string{"SEC Excluded", strconv.Itoa(d.FlagSummary.SecExcluded)})
	writer.Write([]string{"Manual Review", strconv.Itoa(d.FlagSummary.ManualReview)})

	writer.Flush()

	if err := writer.Error(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func generateDashboardXLSX(d *securitisation.DashboardEligibleAccountResponse) ([]byte, error) {
	f := excelize.NewFile()
	sheet := "Dashboard"
	f.NewSheet(sheet)

	// Header
	f.SetCellValue(sheet, "A1", "Metric")
	f.SetCellValue(sheet, "B1", "Value")

	// Rows
	data := [][]string{
		{"Total Loans", strconv.FormatInt(d.TotalLoans, 10)},
		{"Eligible Loans", strconv.FormatInt(d.EligibleLoans, 10)},
		{"Ineligible Loans", strconv.FormatInt(d.IneligibleLoans, 10)},
		{"Total Outstanding", strconv.FormatFloat(d.TotalOutstanding, 'f', 2, 64)},
		{"DD Pass", strconv.Itoa(d.DueDiligenceSummary.Pass)},
		{"DD Fail", strconv.Itoa(d.DueDiligenceSummary.Fail)},
		{"DD Pending", strconv.Itoa(d.DueDiligenceSummary.Pending)},
		{"SEC Ready", strconv.Itoa(d.FlagSummary.SecReady)},
		{"SEC Excluded", strconv.Itoa(d.FlagSummary.SecExcluded)},
		{"Manual Review", strconv.Itoa(d.FlagSummary.ManualReview)},
	}

	row := 2
	for _, item := range data {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), item[0])
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), item[1])
		row++
	}

	// Save to memory
	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
