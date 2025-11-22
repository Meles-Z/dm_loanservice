package duediligence

import (
	"dm_loanservice/drivers/dbmodels"
	duediligence "dm_loanservice/internal/service/domain/due_diligence"

	"github.com/aarondl/null/v8"
)

func mapDueDiligence(dd *dbmodels.Duediligencechecklistitem) duediligence.DueDiligence {
	return duediligence.DueDiligence{
		ID:            dd.ID,
		AccountID:     dd.AccountID,
		ChecklistItem: dd.ChecklistItem,
		Status:        dd.Status,
		Comments:      null.String{String: dd.Comments.String, Valid: dd.Comments.Valid},
		CreatedBy:     null.Int{Int: dd.CreatedBy.Int, Valid: dd.CreatedBy.Valid},
		CreatedAt:     null.Time{Time: dd.CreatedAt.Time, Valid: dd.CreatedAt.Valid},
		UpdatedAt:     null.Time{Time: dd.UpdatedAt.Time, Valid: dd.UpdatedAt.Valid},
	}
}

func calculateOverallStatus(items []*dbmodels.Duediligencechecklistitem) string {
	hasFail := false
	allPass := true

	for _, i := range items {
		switch i.Status {
		case "Fail":
			hasFail = true
			allPass = false
		case "Pending":
			allPass = false
		case "Pass":
			// do nothing
		}
	}

	if hasFail {
		return "Fail"
	}
	if allPass {
		return "Pass"
	}
	return "Pending"
}

func mapDueDiligenceList(items []*dbmodels.Duediligencechecklistitem) []duediligence.DueDiligence {
    mapped := make([]duediligence.DueDiligence, 0, len(items))
    for _, i := range items {
        mapped = append(mapped, mapDueDiligence(i))
    }
    return mapped
}
