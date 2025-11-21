package account

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"dm_loanservice/drivers/dbmodels"
	domain "dm_loanservice/internal/service/domain/account"
	"dm_loanservice/internal/service/domain/dashboard"

	redisLib "dm_loanservice/drivers/redis"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/aarondl/sqlboiler/v4/queries/qm"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

func NewRepository(db *sqlx.DB) domain.Repository {
	return &repository{db: db, schema: "public", redis: redisLib.GetConnection(context.Background())}
}

type repository struct {
	db     *sqlx.DB
	schema string
	redis  *redis.Client
}

func (r *repository) AccountAdd(ctx context.Context, m dbmodels.Account) (*dbmodels.Account, error) {
	err := m.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return &m, nil
}
func (r *repository) AccountId(ctx context.Context, accountID string) (*dbmodels.Account, error) {
	account, err := dbmodels.Accounts(qm.Where("id = ?", accountID)).One(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to get account by ID: %w", err)
	}
	return account, nil
}

func (r *repository) AccountArrears(ctx context.Context) (*dashboard.TotalArrears, error) {
	cacheKey := "dashboard:total_arrears"

	// 1️⃣ Try from Redis first
	cached, err := r.redis.Get(ctx, cacheKey).Result()
	if err == nil && cached != "" {
		var result dashboard.TotalArrears
		if err := json.Unmarshal([]byte(cached), &result); err == nil {
			return &result, nil
		}
	}

	// 2️⃣ Compute live from DB
	count, err := dbmodels.Accounts(qm.Where("arrears_flag = true")).Count(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to count arrears accounts: %w", err)
	}

	var currentValue, prevValue float64

	// Current total arrears value
	if err := r.db.QueryRowContext(ctx, `
		SELECT COALESCE(SUM(arrears_amount), 0)
		FROM accounts
		WHERE arrears_flag = true
	`).Scan(&currentValue); err != nil {
		return nil, fmt.Errorf("failed to calculate current arrears value: %w", err)
	}

	// Previous period arrears (example: last 30 days)
	if err := r.db.QueryRowContext(ctx, `
		SELECT COALESCE(SUM(arrears_amount), 0)
		FROM accounts
		WHERE arrears_flag = true
		AND updated_at < NOW() - INTERVAL '30 days'
	`).Scan(&prevValue); err != nil {
		prevValue = 0 // fallback
	}

	// 3️ Compute trend %
	var trend float64
	if prevValue == 0 {
		trend = 0
	} else {
		trend = ((currentValue - prevValue) / prevValue) * 100
	}

	result := &dashboard.TotalArrears{
		Count: int(count),
		Value: currentValue,
		Trend: trend,
	}

	// 4️⃣ Cache result in Redis (TTL: 5 minutes)
	jsonData, _ := json.Marshal(result)
	if err := r.redis.Set(ctx, cacheKey, jsonData, 5*time.Minute).Err(); err != nil {
		fmt.Printf("⚠️ Redis cache set failed: %v\n", err)
	}

	return result, nil
}

func (r *repository) RecentArrearsCases(ctx context.Context) ([]dashboard.ArrearsCase, error) {
	query := `
	SELECT 
		CONCAT(c.first_name, ' ', c.last_name) AS customer,
		p.referenceid AS mortgageref,
		COALESCE(a.arrears_amount, 0)::TEXT AS arrears,
		TO_CHAR(a.start_date, 'YYYY-MM-DD') AS since
	FROM accounts a
	JOIN customers c ON a.customer_id = c.id
	JOIN products p ON a.product_id = p.id
	WHERE a.arrears_flag = TRUE
	ORDER BY a.start_date DESC
	LIMIT 10;
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()
	var results []dashboard.ArrearsCase
	for rows.Next() {
		var item dashboard.ArrearsCase
		if err := rows.Scan(&item.Customer, &item.MortgageRef, &item.Arrears, &item.Since); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		results = append(results, item)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration failed: %w", err)
	}

	return results, nil
}

func (r *repository) RecentArrears(ctx context.Context) ([]domain.AccountRecentResponse, error) {
	query := `
		SELECT 
			CONCAT(c.first_name, ' ', c.last_name) AS customer_name,
			p.referenceid AS mortgage_id,
			COALESCE(a.arrears_days, 0) AS arrears_age,
			COALESCE(a.arrears_amount, 0) AS amount,
			p.status AS status
		FROM accounts a
		JOIN customers c ON a.customer_id = c.id
		JOIN products p ON a.product_id = p.id
		WHERE a.arrears_flag = TRUE
		ORDER BY a.start_date DESC
		LIMIT 10;
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var results []domain.AccountRecentResponse
	for rows.Next() {
		var item domain.AccountRecentResponse
		if err := rows.Scan(
			&item.CustomerName,
			&item.MortgageID,
			&item.ArrearsAge,
			&item.Amount,
			&item.Status,
		); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		results = append(results, item)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration failed: %w", err)
	}

	return results, nil
}

func (r *repository) MortgagePerformance(ctx context.Context) (*dashboard.MortgagePerformance, error) {
	query := `
		SELECT
			COALESCE(TO_CHAR(COALESCE(new.month, red.month), 'Mon'), '') AS month,
			COALESCE(new.new_mortgages, 0) AS new_mortgages,
			COALESCE(red.redemptions, 0) AS redemptions,
			COALESCE(new.new_mortgages, 0) - COALESCE(red.redemptions, 0) AS net_growth
		FROM
			(SELECT DATE_TRUNC('month', start_date) AS month, COUNT(*) AS new_mortgages
			 FROM accounts
			 GROUP BY 1) AS new
		FULL OUTER JOIN
			(SELECT DATE_TRUNC('month', end_date) AS month, COUNT(*) AS redemptions
			 FROM accounts
			 WHERE end_date <= CURRENT_DATE
			 GROUP BY 1) AS red
		ON new.month = red.month
		ORDER BY COALESCE(new.month, red.month);
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var months []string
	var newMortgages, redemptions, netGrowth []int

	for rows.Next() {
		var (
			month    string
			newCount int
			redCount int
			netCount int
		)
		if err := rows.Scan(&month, &newCount, &redCount, &netCount); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		months = append(months, month)
		newMortgages = append(newMortgages, newCount)
		redemptions = append(redemptions, redCount)
		netGrowth = append(netGrowth, netCount)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration failed: %w", err)
	}

	return &dashboard.MortgagePerformance{
		Months:      months,
		New:         newMortgages,
		Redemptions: redemptions,
		NetGrowth:   netGrowth,
	}, nil
}
