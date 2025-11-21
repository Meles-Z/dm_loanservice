package latefeerule

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
)

type Repository interface {
	LetFeeRuleAdd(context.Context, dbmodels.LateFeeRule) (*dbmodels.LateFeeRule, error)
	LetFeeRuleRead(context.Context, string) (*dbmodels.LateFeeRule, error)
	LetFeeRuleReadByProduct(context.Context, string) ([]*dbmodels.LateFeeRule, error)
	LateFeeRuleUpdate(context.Context, dbmodels.LateFeeRule) (*dbmodels.LateFeeRule, error)
	
}
