package accountflag

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
	"dm_loanservice/drivers/utils"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/drivers/uuid"
	"dm_loanservice/internal/service/domain/account"
	accountauditlog "dm_loanservice/internal/service/domain/account_audit_log"
	accountflag "dm_loanservice/internal/service/domain/account_flag"
	"fmt"
	"time"

	"github.com/aarondl/null/v8"
)

type Service interface {
	AccountFlagAdd(context.Context, *ctxDM.Context, accountflag.AccountFlagAddRequest) (*accountflag.AccountFlagResponse, error)
	AccountFlagRead(context.Context, *ctxDM.Context, accountflag.AccountFlagReadRequest) (*accountflag.AccountFlagResponse, error)
}

func NewService(accountFlagRepo accountflag.Repository, accountRepo account.Repository, accountAuditLog accountauditlog.Repository) Service {
	return &svc{
		accountFlagRepo: accountFlagRepo,
		accountRepo:     accountRepo,
		accountAuditLog: accountAuditLog,
	}
}

type svc struct {
	accountRepo     account.Repository
	accountFlagRepo accountflag.Repository
	accountAuditLog accountauditlog.Repository
}

func (s *svc) AccountFlagAdd(ctx context.Context, ctxDM *ctxDM.Context, req accountflag.AccountFlagAddRequest) (*accountflag.AccountFlagResponse, error) {
	// ✅ AccountID is already assigned from URL path
	if err := req.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		return nil, utils.ErrBadRequest
	}

	acc, err := s.accountRepo.AccountRead(ctx, req.AccountID)
	if err != nil {
		return nil, err
	}

	if acc.FraudFlag.Bool {
		return nil, fmt.Errorf("account flagged as fraud")
	}

	flag := dbmodels.AccountFlag{
		ID:        uuid.UUID(),
		AccountID: req.AccountID,
		FlagType:  req.FlagType,
		Reason:    null.String{String: req.Reason, Valid: true},
		FlaggedBy: int(ctxDM.UserSession.UserId),
	}

	newFlag, err := s.accountFlagRepo.AccountFlagAdd(ctx, flag)
	if err != nil {
		return nil, err
	}

	_, err = s.accountAuditLog.AccountAuditLogAdd(ctx, dbmodels.AccountAuditLog{
		ID:          uuid.UUID(),
		AccountID:   req.AccountID,
		Action:      null.String{String: "Account flagged", Valid: true},
		Details:     null.String{String: fmt.Sprintf("Flag type: %s, Reason: %s", req.FlagType, req.Reason), Valid: true},
		PerformedBy: null.Int{Int: int(ctxDM.UserSession.UserId), Valid: true},
		PerformedAt: null.Time{Time: time.Now()},
	})
	if err != nil {
		return nil, err
	}

	mappedFlag := mapAccountFlag(newFlag)
	return &accountflag.AccountFlagResponse{
		AccountFlag: mappedFlag,
	}, nil
}

func (s *svc) AccountFlagRead(ctx context.Context, ctxDM *ctxDM.Context, req accountflag.AccountFlagReadRequest) (*accountflag.AccountFlagResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, utils.ErrBadRequest
	}
	flag, err := s.accountFlagRepo.AccountFlagRead(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	mappedFlag := mapAccountFlag(flag)
	return &accountflag.AccountFlagResponse{
		AccountFlag: mappedFlag,
	}, nil
}
