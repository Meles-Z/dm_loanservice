package tasks

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
	"dm_loanservice/internal/service/domain/dashboard"
)

type Repository interface {
	TaskAdd(ctx context.Context, inq dbmodels.Task) (*dbmodels.Task, error)
	TaskRead(ctx context.Context, id string) (*dbmodels.Task, error)
	TaskHub(ctx context.Context) ([]dashboard.TaskItem, error)
	RecentTasks(ctx context.Context) (*dbmodels.TaskSlice, error)
	TaskSummary(ctx context.Context) (*TaskSummaryResponse, error)
}
