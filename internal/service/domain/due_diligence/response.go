package duediligence

type DueDiligenceResponse struct {
	DueDiligence DueDiligence `json:"due_diligence"`
}

type DueDiligenceByAccountResponse struct {
	DueDiligence []DueDiligence `json:"due_diligence"`
}

type DueDiligenceUpdateResponse struct {
	OverallStatus string                      `json:"overall_status"`
	Items         []DueDiligence `json:"items"`
}
