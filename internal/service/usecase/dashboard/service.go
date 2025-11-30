package dashboard

import (
	"context"
	iniquiriespb "dm_loanservice/drivers/customer_inquiriesservice/inquiriespb"
	"dm_loanservice/drivers/productservice/productpb"
	ctxDM "dm_loanservice/drivers/utils/context"
	"dm_loanservice/internal/service/domain/account"
	domain "dm_loanservice/internal/service/domain/dashboard"
	"dm_loanservice/internal/service/domain/tasks"
)

type Service interface {
	PortfolioSummary(ctx context.Context, ctxDM *ctxDM.Context) (*domain.DashboardResponse, error)
}

func NewService(accountRepo account.Repository,
	inquiriesRepo iniquiriespb.CustomerInquiriesWrapper,
	task tasks.Repository,
	product productpb.ProductWrapper,
) Service {
	return &svc{
		a:         accountRepo,
		inquiries: inquiriesRepo,
		task:      task,
		product:   product,
	}
}

type svc struct {
	a         account.Repository
	inquiries iniquiriespb.CustomerInquiriesWrapper
	task      tasks.Repository
	product   productpb.ProductWrapper
}
