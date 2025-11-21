package latefeerule

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
	ctxDM "dm_loanservice/drivers/utils/context"
	latefeerule "dm_loanservice/internal/service/domain/late_fee_rule"
)

func (s *svc) LateFeeRuleAdd(ctx context.Context, ctxSess *ctxDM.Context, req latefeerule.LateFeeRuleAddRequest) (*latefeerule.LateFeeRuleAddResponse, error) {
	rule, err := s.lateFeeRuleRepo.LetFeeRuleAdd(ctx, dbmodels.LateFeeRule{
		ProductID:       req.ProductId,
		FeeType:         req.FeeType,
		RateOrAmount:    req.RateOrAmount,
		GracePeriodDays: req.GracePeriodDays,
		MaxFeeAmount:    req.MaxFeeAmount,
		InterestRate:    req.InterestRate,
		RegulatoryCap:   req.RegulatoryCap,
		EffectiveFrom:   req.EffectiveFrom,
		EffectiveTo:     req.EffectiveTo,
	})
	if err != nil {
		return nil, err
	}
	lateFee := mapLateFeeRule(rule)
	return &latefeerule.LateFeeRuleAddResponse{
		LateFeeRule: &lateFee,
	}, nil
}

func (s *svc) LateFeeRuleRead(ctx context.Context, ctxSess *ctxDM.Context, req latefeerule.LateFeeRuleReadRequest) (*latefeerule.LateFeeRuleReadResponse, error) {
	rule, err := s.lateFeeRuleRepo.LetFeeRuleRead(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	lateFee := mapLateFeeRule(rule)
	return &latefeerule.LateFeeRuleReadResponse{
		LateFeeRule: &lateFee,
	}, nil
}

func (s *svc) LateFeeRuleUpdate(ctx context.Context, ctxSess *ctxDM.Context, req latefeerule.LateFeeRuleUpdateRequest) (*latefeerule.LateFeeRuleUpdateResponse, error) {
	rule, err := s.lateFeeRuleRepo.LateFeeRuleUpdate(ctx, dbmodels.LateFeeRule{
		ID:              req.Id,
		ProductID:       req.ProductId,
		FeeType:         req.FeeType,
		RateOrAmount:    req.RateOrAmount,
		GracePeriodDays: req.GracePeriodDays,
		MaxFeeAmount:    req.MaxFeeAmount,
		InterestRate:    req.InterestRate,
		RegulatoryCap:   req.RegulatoryCap,
		EffectiveFrom:   req.EffectiveFrom,
		EffectiveTo:     req.EffectiveTo,
	})
	if err != nil {
		return nil, err
	}
	lateFee := mapLateFeeRule(rule)
	return &latefeerule.LateFeeRuleUpdateResponse{
		LateFeeRule: &lateFee,
	}, nil
}
