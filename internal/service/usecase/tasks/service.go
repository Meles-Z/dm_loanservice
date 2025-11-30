package tasks

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/internal/service/domain/tasks"
)

type Service interface {
	TaskAdd(context.Context, *ctxDM.Context, tasks.TaskAddRequest) (*tasks.TaskAddResponse, error)
	// TaskRead(context.Context, *ctxDM.Context, tasks.tasksReadRequest) (*tasks.TasksReadResponse, error)
	RecentTasks(ctx context.Context, ctxDM *ctxDM.Context) (*dbmodels.TaskSlice, error)
	TaskSummary(ctx context.Context, ctxDM *ctxDM.Context) (*tasks.TaskSummaryResponse, error)
}

func NewService(repo tasks.Repository) Service {
	return &svc{
		r: repo,
	}
}

type svc struct {
	r tasks.Repository
}
