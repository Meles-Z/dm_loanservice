package duediligence

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
	"dm_loanservice/drivers/utils"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/drivers/uuid"
	"dm_loanservice/internal/service/domain/account"
	accountauditlog "dm_loanservice/internal/service/domain/account_audit_log"
	accountflag "dm_loanservice/internal/service/domain/account_flag"
	duediligence "dm_loanservice/internal/service/domain/due_diligence"
	"fmt"
	"time"

	"github.com/aarondl/null/v8"
)

type Service interface {
	DueDiligenceAdd(context.Context, *ctxDM.Context, duediligence.DueDiligenceAddRequest) (*duediligence.DueDiligenceResponse, error)
	DueDiligenceRead(context.Context, *ctxDM.Context, duediligence.DueDiligenceReadRequest) (*duediligence.DueDiligenceResponse, error)
	DueDiligenceUpdate(context.Context, *ctxDM.Context, duediligence.DueDiligenceUpdateRequest) (*duediligence.DueDiligenceUpdateResponse, error)
	DueDiligenceByAccount(context.Context, *ctxDM.Context, duediligence.DueDiligenceByAccountRequest) (*duediligence.DueDiligenceByAccountResponse, error)
}

func NewService(duediligenceRepo duediligence.Repository, accountRepo account.Repository,
	accFlag accountflag.Repository, accountAuditLog accountauditlog.Repository) Service {
	return &svc{
		dueDiligenceRepo: duediligenceRepo,
		accountRepo:      accountRepo,
		accFlagRepo:      accFlag,
		accountAuditLog:  accountAuditLog,
	}
}

type svc struct {
	accountRepo      account.Repository
	dueDiligenceRepo duediligence.Repository
	accFlagRepo      accountflag.Repository
	accountAuditLog  accountauditlog.Repository
}

func (s *svc) DueDiligenceAdd(ctx context.Context, ctxDM *ctxDM.Context, req duediligence.DueDiligenceAddRequest) (*duediligence.DueDiligenceResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, utils.ErrBadRequest
	}
	acc, err := s.accountRepo.AccountRead(ctx, req.AccountID)
	if err != nil {
		ctxDM.ErrorMessage = err.Error()
		return nil, err
	}

	if acc.FraudFlag.Bool {
		return nil, err
	}

	dd := dbmodels.Duediligencechecklistitem{
		ID:            uuid.UUID(),
		AccountID:     req.AccountID,
		ChecklistItem: req.ChecklistItem,
		Status:        req.Status,
		Comments:      null.String{String: req.Comments, Valid: true},
		CreatedBy:     null.Int{Int: int(ctxDM.UserSession.UserId), Valid: true},
	}
	newDD, err := s.dueDiligenceRepo.DueDiligenceAdd(ctx, dd)
	if err != nil {
		return nil, err
	}
	mappedDD := mapDueDiligence(newDD)
	return &duediligence.DueDiligenceResponse{
		DueDiligence: mappedDD,
	}, nil
}

func (s *svc) DueDiligenceRead(ctx context.Context, ctxDM *ctxDM.Context, req duediligence.DueDiligenceReadRequest) (*duediligence.DueDiligenceResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, utils.ErrBadRequest
	}
	dd, err := s.dueDiligenceRepo.DueDiligenceRead(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	mappedDD := mapDueDiligence(dd)
	return &duediligence.DueDiligenceResponse{
		DueDiligence: mappedDD,
	}, nil
}

func (s *svc) DueDiligenceUpdate(
	ctx context.Context,
	ctxDM *ctxDM.Context,
	req duediligence.DueDiligenceUpdateRequest,
) (*duediligence.DueDiligenceUpdateResponse, error) {

	// 1️⃣ Load all checklist items for this account
	ddItems, err := s.dueDiligenceRepo.DueDiligenceByAccount(ctx, req.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch existing DD items: %w", err)
	}

	// Account must have DD items created
	if len(ddItems) == 0 {
		return nil, fmt.Errorf("no due diligence items found for account %s", req.ID)
	}

	// 2️⃣ Identify which item must be updated
	if req.ChecklistItem == nil {
		return nil, fmt.Errorf("checklist_item is required")
	}

	var existingItem *dbmodels.Duediligencechecklistitem
	for _, item := range ddItems {
		if item.ChecklistItem == *req.ChecklistItem {
			existingItem = item
			break
		}
	}
	if existingItem == nil {
		return nil, fmt.Errorf("checklist item '%s' not found for account '%s'",
			*req.ChecklistItem, req.ID)
	}

	// 3️⃣ Update the specific checklist item
	_, err = s.dueDiligenceRepo.DueDiligenceUpdate(ctx, req, existingItem)
	if err != nil {
		return nil, fmt.Errorf("failed to update checklist item: %w", err)
	}

	// Refresh all items after update
	ddItems, err = s.dueDiligenceRepo.DueDiligenceByAccount(ctx, req.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to reload updated items: %w", err)
	}

	// 4️⃣ Recalculate overall DD status
	overallStatus := calculateOverallStatus(ddItems)

	// 5️⃣ Auto-flag if overall = FAIL
	if overallStatus == "Fail" {
		_, err := s.accFlagRepo.AccountFlagAdd(ctx, dbmodels.AccountFlag{
			ID:        uuid.UUID(),
			AccountID: req.ID,
			FlagType:  "due_diligence",
			Reason:    null.String{String: "Due diligence auto-fail", Valid: true},
			FlaggedBy: int(ctxDM.UserSession.UserId),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create account flag: %w", err)
		}
	}

	_, err = s.accountAuditLog.AccountAuditLogAdd(ctx, dbmodels.AccountAuditLog{
		ID:          uuid.UUID(),
		AccountID:   req.ID,
		Action:      null.String{String: "Account flagged", Valid: true},
		Details:     null.String{String: "Due diligence auto-fail", Valid: true},
		PerformedBy: null.Int{Int: int(ctxDM.UserSession.UserId), Valid: true},
		PerformedAt: null.Time{Time: time.Now()},
	})
	if err != nil {
		return nil, err
	}

	// 6️⃣ Build API response with all items + overall status
	mappedItems := mapDueDiligenceList(ddItems)

	return &duediligence.DueDiligenceUpdateResponse{
		OverallStatus: overallStatus,
		Items:         mappedItems,
	}, nil

}

func (s *svc) DueDiligenceByAccount(ctx context.Context, ctxDM *ctxDM.Context, req duediligence.DueDiligenceByAccountRequest) (*duediligence.DueDiligenceByAccountResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, utils.ErrBadRequest
	}
	dd, err := s.dueDiligenceRepo.DueDiligenceByAccount(ctx, req.AccountID)
	if err != nil {
		ctxDM.ErrorMessage = err.Error()
		return nil, err
	}
	mappedDD := make([]duediligence.DueDiligence, len(dd))
	for i, d := range dd {
		mappedDD[i] = mapDueDiligence(d)
	}
	return &duediligence.DueDiligenceByAccountResponse{
		DueDiligence: mappedDD,
	}, nil
}
