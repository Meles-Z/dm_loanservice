package tasks

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
	"dm_loanservice/drivers/utils"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/drivers/uuid"
	"dm_loanservice/internal/service/domain/tasks"
	"time"

	"github.com/aarondl/null/v8"
)

func (s *svc) TaskAdd(ctx context.Context, ctxDM *ctxDM.Context, req tasks.TaskAddRequest) (*tasks.TaskAddResponse, error) {
	if err := req.Validate(); err != nil {
		ctxDM.ErrorMessage = err.Error()
		return nil, utils.ErrorNewInvalidRequest
	}
	parsedTask := dbmodels.Task{
		ID:                  uuid.UUID(),
		CustomerInquiriesID: null.String{String: req.InquiriesID, Valid: true},
		CustomerID:          req.CustomerID,
		AssignedTo:          null.Int{Int: req.AssignedTo, Valid: true},
		Title:               req.Title,
		Description:         null.String{String: req.Description, Valid: true},
		Status:              req.Status,
		DueDate:             null.Time{Time: req.DueDate, Valid: !req.DueDate.IsZero()},
		Priority:            req.Priority,
		CreatedBy:           req.CreatedBy,
	}

	taskAdd, err := s.r.TaskAdd(ctx, parsedTask)
	if err != nil {
		return nil, err
	}

	var updatedAt, deletedAt time.Time
	if taskAdd.UpdatedAt.Valid {
		updatedAt = taskAdd.UpdatedAt.Time
	}
	if taskAdd.DeletedAt.Valid {
		deletedAt = taskAdd.DeletedAt.Time
	}

	taskRes := tasks.Tasks{
		ID:          taskAdd.ID,
		InquiriesID: taskAdd.CustomerInquiriesID.String,
		CustomerID:  taskAdd.CustomerID,
		AssignedTo:  taskAdd.AssignedTo.Int,
		Title:       taskAdd.Title,
		Description: taskAdd.Description.String,
		Status:      taskAdd.Status,
		DueDate:     taskAdd.DueDate.Time,
		Priority:    taskAdd.Priority,
		CreatedAt:   taskAdd.CreatedAt.Time,
		CreatedBy:   taskAdd.CreatedBy,
		UpdatedAt:   updatedAt,
		DeletedAt:   &deletedAt,
	}

	return &tasks.TaskAddResponse{
		Tasks: taskRes,
	}, nil
}
func (s *svc) RecentTasks(ctx context.Context, ctxDM *ctxDM.Context) (*dbmodels.TaskSlice, error) {
	tasks, err := s.r.RecentTasks(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *svc) TaskSummary(ctx context.Context, ctxDM *ctxDM.Context) (*tasks.TaskSummaryResponse, error) {
	counts, err := s.r.TaskSummary(ctx)
	if err != nil {
		return nil, err
	}
	return counts, nil
}
