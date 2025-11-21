package accountflag

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/drivers/uuid"
	"dm_loanservice/internal/service/domain/account"
	accountflag "dm_loanservice/internal/service/domain/account_flag"

	"github.com/aarondl/null/v8"
)

type Service interface {
	AccountFlagAdd(context.Context, *ctxDM.Context, accountflag.AccountFlagAddRequest) (*accountflag.AccountFlagResponse, error)
	AccountFlagRead(context.Context, *ctxDM.Context, accountflag.AccountFlagReadRequest) (*accountflag.AccountFlagResponse, error)
}

func NewService(accountFlagRepo accountflag.Repository, accountRepo account.Repository) Service {
	return &svc{
		accountFlagRepo: accountFlagRepo,
		accountRepo:     accountRepo,
	}
}

type svc struct {
	accountRepo     account.Repository
	accountFlagRepo accountflag.Repository
}

func (s *svc) AccountFlagAdd(ctx context.Context, ctxDM *ctxDM.Context, req accountflag.AccountFlagAddRequest) (*accountflag.AccountFlagResponse, error) {
	acc, err := s.accountRepo.AccountRead(ctx, req.AccountID)
	if err != nil {
		ctxDM.ErrorMessage = err.Error()
		return nil, err
	}

	if acc.FraudFlag.Bool {
		return nil, err
	}

	flag := dbmodels.AccountFlag{
		ID:        uuid.UUID(),
		AccountID: req.AccountID,
		FlagType:  null.String{String: req.FlagType, Valid: true},
		Reason:    null.String{String: req.Reason, Valid: true},
		FlaggedBy: null.Int{Int: int(ctxDM.UserSession.UserId), Valid: true},
	}
	newFlag, err := s.accountFlagRepo.AccountFlagAdd(ctx, flag)
	if err != nil {
		return nil, err
	}
	mappedFlag := mapAccountFlag(newFlag)
	return &accountflag.AccountFlagResponse{
		AccountFlag: mappedFlag,
	}, nil
}

func (s *svc) AccountFlagRead(ctx context.Context, ctxDM *ctxDM.Context, req accountflag.AccountFlagReadRequest) (*accountflag.AccountFlagResponse, error) {
	flag, err := s.accountFlagRepo.AccountFlagRead(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	mappedFlag := mapAccountFlag(flag)
	return &accountflag.AccountFlagResponse{
		AccountFlag: mappedFlag,
	}, nil
}
