package servicerestriction

import (
	"dm_loanservice/drivers/dbmodels"
	servicerestriction "dm_loanservice/internal/service/domain/service_restriction"
)

func mapServiceRestriction(serviceRestriction *dbmodels.ServicingRestriction) *servicerestriction.ServiceRestriction {
	return &servicerestriction.ServiceRestriction{
		ID:              serviceRestriction.ID,
		AccountID:       serviceRestriction.AccountID.String,
		RestrictionType: serviceRestriction.RestrictionType,
		ActionName:      serviceRestriction.ActionName,
		IsActive:        serviceRestriction.IsActive,
		Reason:          serviceRestriction.Reason,
		Source:          serviceRestriction.Source.String,
		CreatedAt:       serviceRestriction.CreatedAt.String(),
		UpdatedAt:       serviceRestriction.UpdatedAt.String(),
	}
}
