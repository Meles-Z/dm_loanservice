package tasks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

type ArrearsHandler interface {
	RecalculateArrears(ctx context.Context) error
}

type Processor struct {
	handler ArrearsHandler
}

func NewProcessor(handler ArrearsHandler) *Processor {
	return &Processor{handler: handler}
}

func (p *Processor) Register(mux *asynq.ServeMux) {
	mux.HandleFunc(TaskCalculateArrears, p.handleArrearsCalculation)
}

func (p *Processor) handleArrearsCalculation(ctx context.Context, t *asynq.Task) error {
	var payload ArrearsPayload
	if t.Payload() != nil && len(t.Payload()) > 0 {
		if err := json.Unmarshal(t.Payload(), &payload); err != nil {
			fmt.Println("⚠️ Warning: invalid payload, continuing:", err)
		}
	} else {
		payload.RunType = "scheduled" // default value
	}

	fmt.Println("🔄 Running arrears recalculation job:", payload.RunType)

	if err := p.handler.RecalculateArrears(ctx); err != nil {
		fmt.Println("❌ RecalculateArrears failed:", err)
		return err
	}

	fmt.Println("✅ Arrears calculation completed successfully!")
	return nil
}
