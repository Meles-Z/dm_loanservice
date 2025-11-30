package grpc

import (
	"context"
	"dm_loanservice/internal/service/domain/account"

	"github.com/aarondl/sqlboiler/v4/types"
	"github.com/brianjobling/dm_proto/generated/accountservice/accountpb"
	"github.com/ericlagergren/decimal"
)

func decodeAccountAddReq(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*accountpb.AccountCreateReq)

	startDate, err := parseDateFlexible(req.StartDate)
	if err != nil {
		return nil, err
	}

	endDate, err := parseDateFlexible(req.EndDate)
	if err != nil {
		return nil, err
	}

	loanBig := new(decimal.Big).SetFloat64(req.LoanAmount)
	balanceBig := new(decimal.Big).SetFloat64(req.BalanceOutstanding)
	arreasAmount := new(decimal.Big).SetFloat64(req.ArrearsAmount)
	return account.AccountAddRequest{
		CustomerID:         req.CustomerId,
		ProductID:          req.ProductId,
		LoanAmount:         types.NewDecimal(loanBig),
		BalanceOutstanding: types.NewDecimal(balanceBig),
		StartDate:          startDate,
		EndDate:            endDate,
		TermYears:          int(req.TermYears),

		ArrearsFlag:   req.ArrearsFlag,
		ArrearsAmount: types.NewDecimal(arreasAmount),
		ArrearsDays:   int(req.ArrearsDays),

		ForbearanceFlag: req.ForbearanceFlag,
		ForbearanceType: account.ForbearanceType(req.ForbearanceType),

		FraudFlag:  req.FraudFlag,
		FraudNotes: req.FraudNotes,

		RedrawFacility:    req.RedrawFacility,
		CollateralAddress: req.CollateralAddress,
		CollateralType:    account.CollateralType(req.CollateralType),
		SecurityType:      account.SecurityType(req.SecurityType),

		PortfolioID:         req.PortfolioId,
		StressTestResult:    account.StressTestResult(req.StressTestResult),
		CapitalAdequacyFlag: req.CapitalAdequacyFlag,
	}, nil
}

func decodeAccountReadReq(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*accountpb.AccountGetReq)
	return account.AccountReadRequest{
		ID: req.Id,
	}, nil
}
