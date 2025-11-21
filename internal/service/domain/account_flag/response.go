package accountflag

type (
	AccountFlagResponse struct {
		AccountFlag AccountFlag `json:"account_flag"`
	}
	AccountFlagByAccountResponse struct {
		AccountFlag []AccountFlag `json:"account_flag"`
	}
)
