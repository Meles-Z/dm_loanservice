package latefeerule

import (
	"context"
	latefeerule "dm_loanservice/internal/service/domain/late_fee_rule"
	ctxDM "dm_loanservice/drivers/utils/context"
)

type Service interface {
	LateFeeRuleAdd(ctx context.Context, ctxSess *ctxDM.Context, req latefeerule.LateFeeRuleAddRequest) (*latefeerule.LateFeeRuleAddResponse, error)
	LateFeeRuleRead(ctx context.Context, ctxSess *ctxDM.Context, req latefeerule.LateFeeRuleReadRequest) (*latefeerule.LateFeeRuleReadResponse, error)
	LateFeeRuleUpdate(ctx context.Context, ctxSess *ctxDM.Context, req latefeerule.LateFeeRuleUpdateRequest) (*latefeerule.LateFeeRuleUpdateResponse, error)
}

func NewService(lateFeeRuleRepo latefeerule.Repository) Service {
	return &svc{
		lateFeeRuleRepo: lateFeeRuleRepo,
	}
}

type svc struct {
	lateFeeRuleRepo latefeerule.Repository
}
