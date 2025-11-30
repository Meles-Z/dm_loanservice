package productpb

import (
	ctxDM "dm_loanservice/drivers/utils/context"

	productPb "github.com/brianjobling/dm_proto/generated/productservice/productpb"
)

type ProductWrapper interface {
	ProductRead(ctxSess *ctxDM.Context, id string) (*productPb.ProductReadRes, error)
	ProductByCustomerID(ctxSess *ctxDM.Context, id string) (*productPb.FindProductsByCustomerIDRes, error)
	ProductOverviews(ctxSess *ctxDM.Context) (*productPb.ProductOverviewRes, error)
}
