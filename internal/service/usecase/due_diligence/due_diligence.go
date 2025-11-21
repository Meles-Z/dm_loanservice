package duediligence

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/drivers/uuid"
	"dm_loanservice/internal/service/domain/account"
	duediligence "dm_loanservice/internal/service/domain/due_diligence"
	"fmt"

	"github.com/aarondl/null/v8"
)

type Service interface {
	DueDiligenceAdd(context.Context, *ctxDM.Context, duediligence.DueDiligenceAddRequest) (*duediligence.DueDiligenceResponse, error)
	DueDiligenceRead(context.Context, *ctxDM.Context, duediligence.DueDiligenceReadRequest) (*duediligence.DueDiligenceResponse, error)
	DueDiligenceUpdate(context.Context, *ctxDM.Context, duediligence.DueDiligenceUpdateRequest) (*duediligence.DueDiligenceResponse, error)
	DueDiligenceByAccount(context.Context, *ctxDM.Context, duediligence.DueDiligenceByAccountRequest) (*duediligence.DueDiligenceByAccountResponse, error)
}

func NewService(duediligenceRepo duediligence.Repository, accountRepo account.Repository) Service {
	return &svc{
		dueDiligenceRepo: duediligenceRepo,
		accountRepo:      accountRepo,
	}
}

type svc struct {
	accountRepo      account.Repository
	dueDiligenceRepo duediligence.Repository
}

func (s *svc) DueDiligenceAdd(ctx context.Context, ctxDM *ctxDM.Context, req duediligence.DueDiligenceAddRequest) (*duediligence.DueDiligenceResponse, error) {
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
	dd, err := s.dueDiligenceRepo.DueDiligenceRead(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	mappedDD := mapDueDiligence(dd)
	return &duediligence.DueDiligenceResponse{
		DueDiligence: mappedDD,
	}, nil
}

func (s *svc) DueDiligenceUpdate(ctx context.Context, ctxDM *ctxDM.Context, req duediligence.DueDiligenceUpdateRequest) (*duediligence.DueDiligenceResponse, error) {
	
	fmt.Println("Due diligence service is here")

	// 1️⃣ Fetch all checklist items for the account
	ddItems, err := s.dueDiligenceRepo.DueDiligenceByAccount(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	// 2️⃣ Find the checklist item to update
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
		return nil, fmt.Errorf("checklist item %s not found for account %s", *req.ChecklistItem, req.ID)
	}

	// 3️⃣ Update the checklist item
	_, err = s.dueDiligenceRepo.DueDiligenceUpdate(ctx, req, existingItem)
	if err != nil {
		return nil, err
	}

	// 4️ Recalculate overall status for the account
	// overallStatus := calculateOverallStatus(ddItems)

	// 5️⃣ Map response
	// mappedDD := mapDueDiligenceList(ddItems, overallStatus) // return all items with updated status

	return &duediligence.DueDiligenceResponse{
		// DueDiligence: mappedDD.Items,
	}, nil
}

func (s *svc) DueDiligenceByAccount(ctx context.Context, ctxDM *ctxDM.Context, req duediligence.DueDiligenceByAccountRequest) (*duediligence.DueDiligenceByAccountResponse, error) {
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
