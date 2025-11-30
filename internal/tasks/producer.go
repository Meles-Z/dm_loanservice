package tasks

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
)

type ArrearsPayload struct {
	RunType string `json:"run_type"` // e.g. "nightly" or "manual"
}

type Producer struct {
	client *asynq.Client
}

func NewProducer(redisAddr string) *Producer {
	return &Producer{
		client: asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr}),
	}
}

func (p *Producer) EnqueueArrearsCalculation(ctx context.Context, runType string) error {
	payload, _ := json.Marshal(ArrearsPayload{RunType: runType})
	task := asynq.NewTask(TaskCalculateArrears, payload)
	_, err := p.client.EnqueueContext(ctx, task, asynq.Queue("arrears"), asynq.MaxRetry(3))
	return err
}
