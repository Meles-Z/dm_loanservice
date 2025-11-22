package accountauditlog

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/drivers/uuid"
	"dm_loanservice/internal/service/domain/account"
	accountauditlog "dm_loanservice/internal/service/domain/account_audit_log"
	"fmt"
	"time"

	"github.com/aarondl/null/v8"
)

type Service interface {
	AccountAuditLogAdd(context.Context, *ctxDM.Context, accountauditlog.AccountAuditLogAddRequest) (*accountauditlog.AccountAuditLogResponse, error)
	AccountAuditLogRead(context.Context, *ctxDM.Context, accountauditlog.AccountAuditLogReadRequest) (*accountauditlog.AccountAuditLogResponse, error)
}

func NewService(accountAuditLogRepo accountauditlog.Repository, accountRepo account.Repository) Service {
	return &svc{
		accountAuditLogRepo: accountAuditLogRepo,
		accountRepo:         accountRepo,
	}
}

type svc struct {
	accountAuditLogRepo accountauditlog.Repository
	accountRepo         account.Repository
}

func (s *svc) AccountAuditLogAdd(ctx context.Context, ctxDM *ctxDM.Context, req accountauditlog.AccountAuditLogAddRequest) (*accountauditlog.AccountAuditLogResponse, error) {
	acc, err := s.accountRepo.AccountRead(ctx, req.AccountID)
	if err != nil {
		ctxDM.ErrorMessage = err.Error()
		return nil, err
	}

	if acc.FraudFlag.Bool {
		return nil, fmt.Errorf("account flagged as fraud")
	}

	log := dbmodels.AccountAuditLog{
		ID:          uuid.UUID(),
		AccountID:   req.AccountID,
		Action:      null.String{String: req.Action, Valid: true},
		Details:     null.String{String: req.Details, Valid: true},
		PerformedBy: null.Int{Int: int(ctxDM.UserSession.UserId), Valid: true},
		PerformedAt: null.Time{Time: time.Now(), Valid: true},
	}
	newLog, err := s.accountAuditLogRepo.AccountAuditLogAdd(ctx, log)
	if err != nil {
		return nil, err
	}
	mappedLog := mapAccountAuditLog(newLog)
	return &accountauditlog.AccountAuditLogResponse{
		AccountAuditLog: mappedLog,
	}, nil
}

func (s *svc) AccountAuditLogRead(ctx context.Context, ctxDM *ctxDM.Context, req accountauditlog.AccountAuditLogReadRequest) (*accountauditlog.AccountAuditLogResponse, error) {
	log, err := s.accountAuditLogRepo.AccountAuditLogRead(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	mappedLog := mapAccountAuditLog(log)
	return &accountauditlog.AccountAuditLogResponse{
		AccountAuditLog: mappedLog,
	}, nil
}
