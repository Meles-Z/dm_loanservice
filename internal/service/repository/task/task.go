package task

import (
	"context"
	"dm_loanservice/drivers/dbmodels"
	"dm_loanservice/internal/service/domain/dashboard"
	domain "dm_loanservice/internal/service/domain/tasks"
	"fmt"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/aarondl/sqlboiler/v4/queries/qm"
	"github.com/jmoiron/sqlx"
)

func NewRepository(db *sqlx.DB) domain.Repository {
	return &repository{db: db, schema: "public"}
}

type repository struct {
	db     *sqlx.DB
	schema string
}

func (r *repository) TaskAdd(ctx context.Context, t dbmodels.Task) (*dbmodels.Task, error) {
	err := t.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *repository) TaskRead(ctx context.Context, id string) (*dbmodels.Task, error) {
	return dbmodels.Tasks(qm.Where("id=?", id)).One(ctx, r.db)
}

func (r *repository) TaskHub(ctx context.Context) ([]dashboard.TaskItem, error) {
	query := `
		SELECT 
			COALESCE(u.first_name || ' ' || u.surname, 'Unknown') AS agent,
			t.title AS task,
			t.status
		FROM tasks t
		LEFT JOIN users u ON t.assigned_to = u.id
		ORDER BY t.due_date ASC
		LIMIT 10;
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var results []dashboard.TaskItem
	for rows.Next() {
		var item dashboard.TaskItem
		if err := rows.Scan(&item.Agent, &item.Task, &item.Status); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		results = append(results, item)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration failed: %w", err)
	}

	return results, nil
}

func (r *repository) RecentTasks(ctx context.Context) (*dbmodels.TaskSlice, error) {
	tasks, err := dbmodels.Tasks(
		qm.OrderBy("created_at DESC"),
		qm.Limit(10),
	).All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return &tasks, nil
}

func (r *repository) TaskSummary(ctx context.Context) (*domain.TaskSummaryResponse, error) {
	totalCounts, err := dbmodels.Tasks().Count(ctx, r.db)
	if err != nil {
		return nil, err
	}

	pendingTasks, err := dbmodels.Tasks(qm.Where("status", "pending")).Count(ctx, r.db)
	if err != nil {
		return nil, err
	}

	inProgressTasks, err := dbmodels.Tasks(qm.Where("status", "in_progress")).Count(ctx, r.db)
	if err != nil {
		return nil, err
	}
	reviewTasks, err := dbmodels.Tasks(qm.Where("status", "review")).Count(ctx, r.db)
	if err != nil {
		return nil, err
	}
	blockedTasks, err := dbmodels.Tasks(qm.Where("status", "blocked")).Count(ctx, r.db)
	if err != nil {
		return nil, err
	}
	completedTasks, err := dbmodels.Tasks(qm.Where("status", "completed")).Count(ctx, r.db)
	if err != nil {
		return nil, err
	}
	canceldTasks, err := dbmodels.Tasks(qm.Where("status", "cancelled")).Count(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return &domain.TaskSummaryResponse{
		TotalCounts: totalCounts,
		Pending:     pendingTasks,
		InProgress:  inProgressTasks,
		Review:      reviewTasks,
		Blocked:     blockedTasks,
		Completed:   completedTasks,
		Cancelled:   canceldTasks,
	}, nil
}
