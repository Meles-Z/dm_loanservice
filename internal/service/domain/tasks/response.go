package tasks

type TaskAddResponse struct {
	Tasks Tasks `json:"tasks"`
}

type TaskSummaryResponse struct {
	TotalCounts int64 `json:"totalCount"`
	Pending     int64 `json:"pending"`
	InProgress  int64 `json:"in_progress"`
	Review      int64 `json:"review"`
	Blocked     int64 `json:"blocked"`
	Completed   int64 `json:"completed"`
	Cancelled   int64 `json:"cancelled"`
}
