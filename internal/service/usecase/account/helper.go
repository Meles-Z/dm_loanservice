package account

import (
	"dm_loanservice/drivers/dbmodels"
	"dm_loanservice/internal/service/domain/account"

	"github.com/aarondl/sqlboiler/v4/types"
)

func mapAccount(acc *dbmodels.Account) *account.Accounts {
	return &account.Accounts{
		ID:                 acc.ID,
		CustomerID:         acc.CustomerID,
		ProductID:          acc.ProductID,
		LoanAmount:         acc.LoanAmount,
		BalanceOutstanding: acc.BalanceOutstanding,
		StartDate:          acc.StartDate,
		EndDate:            acc.EndDate,
		TermYears:          acc.TermYears,

		ArrearsFlag:   acc.ArrearsFlag.Bool,
		ArrearsAmount: types.NewDecimal(acc.ArrearsAmount.Big),
		ArrearsDays:   acc.ArrearsDays.Int,

		ForbearanceFlag: acc.ForbearanceFlag.Bool,
		ForbearanceType: account.ForbearanceType(acc.ForbearanceType.String),

		FraudFlag:  acc.FraudFlag.Bool,
		FraudNotes: acc.FraudNotes.String,

		RedrawFacility:    acc.RedrawFacility.Bool,
		CollateralAddress: acc.CollateralAddress.String,
		CollateralType:    account.CollateralType(acc.CollateralType.String),
		SecurityType:      account.SecurityType(acc.SecurityType.String),

		PortfolioID:         acc.PortfolioID.String,
		StressTestResult:    account.StressTestResult(acc.StressTestResult.String),
		CapitalAdequacyFlag: acc.CapitalAdequacyFlag.Bool,

		CreatedAt: acc.CreatedAt.Time,
		CreatedBy: acc.CreatedAt.Time.String(),
		UpdatedAt: acc.UpdatedAt.Time,
	}
}
