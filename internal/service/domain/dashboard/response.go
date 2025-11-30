package dashboard

import (
	productPb "github.com/brianjobling/dm_proto/generated/productservice/productpb"
)

type (
	// Root dashboard response
	DashboardResponse struct {
		PortfolioMetrics    PortfolioMetrics             `json:"portfolioMetrics"`
		RecentArrearsCases  []ArrearsCase                `json:"recentArrearsCases"`
		TaskHub             []TaskItem                   `json:"taskHub"`
		RecentSearches      []string                     `json:"recentSearches"`
		MortgagePerformance MortgagePerformance          `json:"mortgagePerformance"`
		ProductOverview     *productPb.ProductOverviewRes `json:"productOverview"`
	}

	// === Portfolio metrics section ===
	PortfolioMetrics struct {
		TotalArrears      TotalArrears      `json:"totalArrears"`
		UpcomingPayments  UpcomingPayments  `json:"upcomingPayments"`
		NewApplications   NewApplication    `json:"newApplications"`
		CustomerInquiries CustomerInquiries `json:"customerInquiries"`
	}

	TotalArrears struct {
		Count int     `json:"count"`
		Value float64 `json:"value"`
		Trend float64 `json:"trend"`
	}

	UpcomingPayments struct {
		Count  int    `json:"count"`
		Window string `json:"window"`
	}

	NewApplication struct {
		Count  int    `json:"count"`
		Status string `json:"status"`
	}

	CustomerInquiries struct {
		Count  int `json:"count"`
		Urgent int `json:"urgent"`
	}

	// === Recent arrears cases section ===
	ArrearsCase struct {
		Customer    string `json:"customer" db:"customer"`
		MortgageRef string `json:"mortgageRef" db:"mortgageref"` // lowercase for binding
		Arrears     string `json:"arrears" db:"arrears"`
		Since       string `json:"since" db:"since"`
	}

	// === Task hub section ===
	TaskItem struct {
		Agent  string `json:"agent"`
		Task   string `json:"task"`
		Status string `json:"status"`
	}

	// === Mortgage performance section ===
	MortgagePerformance struct {
		Months      []string `json:"months"`
		New         []int    `json:"new"`
		Redemptions []int    `json:"redemptions"`
		NetGrowth   []int    `json:"netGrowth"`
	}

	// === Product overview section ===
	ProductOverview struct {
		Fixed    int64 `json:"Fixed"`
		Variable int64 `json:"Variable"`
		Tracker  int64 `json:"Tracker"`
		Other    int64 `json:"Other"`
	}
)

/*

{
  "portfolioMetrics": {
    "totalArrears": { "count": 27, "value": 204475.0, "trend": -10 },
    "upcomingPayments": { "count": 12, "window": "48h" },
    "newApplications": { "count": 8, "status": "pending_review" },
    "customerInquiries": { "count": 15, "urgent": 5 }
  },
  "recentArrearsCases": [
    { "customer": "James Davies", "mortgageRef": "MOR-29382", "arrears": "£1,149.60", "since": "2023-04-12" }
  ],
  "taskHub": [
    { "agent": "James Davies", "task": "Compliance review", "status": "complete" },
    { "agent": "Robert Harris", "task": "Arrears escalated", "status": "overdue" }
  ],
  "recentSearches": [
    "MOR-29384 (James Davies)", "Fixed Rate 5-Year Products", "12 Oakwood Drive"
  ],
  "mortgagePerformance": {
    "months": ["Jan", "Feb", "Mar", "Apr"],
    "new": [100, 120, 150, 130],
    "redemptions": [50, 60, 70, 65],
    "netGrowth": [50, 60, 80, 65]
  },
  "productOverview": {
    "Fixed": 60,
    "Variable": 25,
    "Tracker": 10,
    "Other": 5
  }
}

*/
