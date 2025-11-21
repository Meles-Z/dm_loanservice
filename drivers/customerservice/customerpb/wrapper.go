package customerpb

import (
	ctxDM "dm_loanservice/drivers/utils/context" // dm_productservice/drivers/utils/context
	"github.com/brianjobling/dm_proto/generated/userservice/customerpb"

)

type CustomerWrapper interface {
	FindByProductID(ctxSess *ctxDM.Context, productID string) ([]*customerpb.Customers, error)
}
