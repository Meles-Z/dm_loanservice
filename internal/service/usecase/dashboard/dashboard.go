package dashboard

import (
	"context"
	ctxDM "dm_loanservice/drivers/utils/context"
	domain "dm_loanservice/internal/service/domain/dashboard"
	"fmt"
	"strconv"
)

func (s *svc) PortfolioSummary(
	ctx context.Context,
	ctxDM *ctxDM.Context,
) (*domain.DashboardResponse, error) {

	fmt.Println("We have reached here")

	fmt.Println("PortfolioSummary called")
	if s.a == nil {
		panic("account repo is nil")
	}
	if s.inquiries == nil {
		panic("inquiries wrapper is nil")
	}
	if s.product == nil {
		panic("product wrapper is nil")
	}
	if s.task == nil {
		panic("task repo is nil")
	}
	// --- Portfolio metrics ---
	totalArrears, err := s.a.AccountArrears(ctx)
	if err != nil {
		return nil, err
	}

	newApplication, err := s.inquiries.CustomerInquiriesPendingCount(ctxDM)
	if err != nil {
		return nil, err
	}

	customerInquiries, err := s.inquiries.CustomerInquiriesUrgentCount(ctxDM)
	if err != nil {
		return nil, err
	}

	// --- Other dashboard sections ---
	recentArrearsCase, err := s.a.RecentArrearsCases(ctx)
	if err != nil {
		return nil, err
	}

	taskHub, err := s.task.TaskHub(ctx)
	if err != nil {
		return nil, err
	}

	mortgagePerformance, err := s.a.MortgagePerformance(ctx)
	if err != nil {
		return nil, err
	}

	productOverview, err := s.product.ProductOverviews(ctxDM)
	if err != nil {
		return nil, err
	}
	urgent, err := strconv.Atoi(customerInquiries.Urgent)
	if err != nil {
		return nil, err
	}

	// --- Compose Portfolio Metrics ---
	portfolioMetrics := domain.PortfolioMetrics{
		TotalArrears: *totalArrears,
		UpcomingPayments: domain.UpcomingPayments{
			Count:  0,
			Window: "",
		},
		NewApplications: domain.NewApplication{
			Count:  int(newApplication.Count),
			Status: newApplication.Status,
		},
		CustomerInquiries: domain.CustomerInquiries{
			Count:  int(customerInquiries.Count),
			Urgent: urgent,
		},
	}

	// --- Final Response ---
	resp := &domain.DashboardResponse{
		PortfolioMetrics:    portfolioMetrics,
		RecentArrearsCases:  recentArrearsCase,
		TaskHub:             taskHub,
		RecentSearches:      []string{}, // or remove if unused
		MortgagePerformance: *mortgagePerformance,
		ProductOverview:     productOverview, // *productPb.ProductOverviewRes
	}

	return resp, nil
}
