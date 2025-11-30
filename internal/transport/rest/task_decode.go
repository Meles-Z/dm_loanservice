package rest

import (
	"context"
	"dm_loanservice/internal/service/domain/tasks"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func decodeTaskAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		err error
		req tasks.TaskAddRequest
	)

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.WithMessage(err, "error decode request")
	}

	return req, nil
}

func decodeRecentTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {

	return nil, nil
}

func decodeTaskSummaryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
