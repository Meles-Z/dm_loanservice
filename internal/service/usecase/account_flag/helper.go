package accountflag

import (
	"dm_loanservice/drivers/dbmodels"
	accountflag "dm_loanservice/internal/service/domain/account_flag"
)

func mapAccountFlag(flag *dbmodels.AccountFlag) accountflag.AccountFlag {
	return accountflag.AccountFlag{
		ID:        flag.ID,
		AccountID: flag.AccountID,
		FlagType:  flag.FlagType,
		Reason:    flag.Reason.String,
		FlaggedBy: flag.FlaggedBy,
		FlaggedAt: flag.FlaggedAt.Time,
	}
}
