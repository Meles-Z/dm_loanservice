package dashboard

// func (s *svc) PortfolioSummary(ctx context.Context, ctxDM *ctxDM.Context) (*domain.DashboardResponse, error) {
// 	totalArrears, err := s.a.AccountArrears(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	newApplication, err := s.inquiries.CustomerInquiriesPendingCount(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	customerInquiries, err := s.inquiries.CustomerInquiriesUrgentCount(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	recentArrearsCase, err := s.a.RecentArrearsCases(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	taskHub, err := s.task.TaskHub(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	mortgagePerformance, err := s.a.MortgagePerformance(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	productOverviews, err := s.product.ProductOverviews(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	portFolioMetrics := domain.PortfolioMetrics{
// 		TotalArrears:      *totalArrears,
// 		NewApplications:   *newApplication,
// 		CustomerInquiries: *customerInquiries,
// 	}
// 	res := &domain.DashboardResponse{
// 		PortfolioMetrics:    portFolioMetrics,
// 		RecentArrearsCases:  recentArrearsCase,
// 		TaskHub:             taskHub,
// 		MortgagePerformance: *mortgagePerformance,
// 		ProductOverview:     *productOverviews,
// 	}
// 	return res, nil
// }
