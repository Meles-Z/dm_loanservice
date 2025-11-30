package iniquiriespb

import (
	ctxDM "dm_loanservice/drivers/utils/context" // dm_productservice/drivers/utils/context

	pb "github.com/brianjobling/dm_proto/generated/customerservice/customer_iniquiriespb"
)

type CustomerInquiriesWrapper interface {
	CustomerInquiriesPendingCount(*ctxDM.Context) (*pb.CustomerInquiriesPendingCountRes, error)
	CustomerInquiriesUrgentCount(*ctxDM.Context) (*pb.CustomerInquiriesUrgentCountRes, error)
}
