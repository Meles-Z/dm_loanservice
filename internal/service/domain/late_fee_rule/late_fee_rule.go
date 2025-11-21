package latefeerule

import (
	"time"

	"github.com/aarondl/sqlboiler/v4/types"
)

type LateFeeRule struct {
	Id              string        `json:"id"`
	ProductId       string        `json:"product_id"`
	FeeType         string        `json:"fee_type"`
	RateOrAmount    types.Decimal `json:"rate_or_amount"`
	GracePeriodDays int           `json:"grace_period_days"`
	MaxFeeAmount    types.Decimal `json:"max_fee_amount"`
	InterestRate    types.Decimal `json:"interest_rate"`
	RegulatoryCap   types.Decimal `json:"regulatory_cap"`
	EffectiveFrom   time.Time     `json:"effective_from"`
	EffectiveTo     time.Time     `json:"effective_to"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}
