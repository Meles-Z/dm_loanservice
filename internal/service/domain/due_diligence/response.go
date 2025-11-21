package duediligence

type DueDiligenceResponse struct {
	DueDiligence DueDiligence `json:"due_diligence"`
}

type DueDiligenceByAccountResponse struct {
	DueDiligence []DueDiligence `json:"due_diligence"`
}
