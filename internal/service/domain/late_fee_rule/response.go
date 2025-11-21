package latefeerule

type (
	LateFeeRuleAddResponse struct {
		LateFeeRule *LateFeeRule `json:"late_fee_rule"`
	}

	LateFeeRuleReadResponse struct {
		LateFeeRule *LateFeeRule `json:"late_fee_rule"`
	}

	LateFeeRuleUpdateResponse struct {
		LateFeeRule *LateFeeRule `json:"late_fee_rule"`
	}
)
