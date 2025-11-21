package latefeerule

import (
	"dm_loanservice/drivers/dbmodels"
	latefeerule "dm_loanservice/internal/service/domain/late_fee_rule"
)

func mapLateFeeRule(req *dbmodels.LateFeeRule) latefeerule.LateFeeRule {
	return latefeerule.LateFeeRule{
		Id:              req.ID,
		ProductId:       req.ProductID,
		FeeType:         req.FeeType,
		RateOrAmount:    req.RateOrAmount,
		GracePeriodDays: req.GracePeriodDays,
		MaxFeeAmount:    req.MaxFeeAmount,
		InterestRate:    req.InterestRate,
		RegulatoryCap:   req.RegulatoryCap,
		EffectiveFrom:   req.EffectiveFrom,
		EffectiveTo:     req.EffectiveTo,
		CreatedAt:       req.CreatedAt,
		UpdatedAt:       req.UpdatedAt,
	}
}
