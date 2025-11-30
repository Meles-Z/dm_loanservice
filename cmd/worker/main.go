package main

import (
	"context"
	"dm_loanservice/drivers/goconf"
	"dm_loanservice/drivers/logger"
	"dm_loanservice/drivers/postgres"
	redisLib "dm_loanservice/drivers/redis"
	accountR "dm_loanservice/internal/service/repository/account"
	accountLockRuleR "dm_loanservice/internal/service/repository/account_lock_rule"
	duediligenceR "dm_loanservice/internal/service/repository/due_diligence"
	serviceRestrictionR "dm_loanservice/internal/service/repository/service_restriction"
	"dm_loanservice/internal/tasks"

	accountSvc "dm_loanservice/internal/service/usecase/account"

	"log"

	"github.com/hibiken/asynq"
)

func main() {
	// Initialize configuration
	cfg := goconf.Config()
	logger.InitLogger()

	logger.LogInfo("🚀 Starting DM UserService Asynq worker...")

	pgdb, err := postgres.NewDBMaster()
	if err != nil {
		logger.LogError("❌ Failed to connect to Postgres: %v", err)
	}
	defer pgdb.Close()

	_ = redisLib.GetConnection(context.Background())
	du := duediligenceR.NewRepository(pgdb)
	lockRule := accountLockRuleR.NewRepository(pgdb)
	serviceRestriction := serviceRestrictionR.NewRepository(pgdb)
	// --- Initialize Service Layer ---
	accountRepo := accountR.NewRepository(pgdb)
	accountService := accountSvc.NewService(accountRepo, du, lockRule, serviceRestriction)

	// --- Asynq Worker Configuration ---
	redisAddr := cfg.GetString("redis.address")
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			Concurrency: 3,
			Queues: map[string]int{
				"arrears": 1, // queue name and priority
			},
		},
	)

	// --- Register Handlers ---
	mux := asynq.NewServeMux()
	processor := tasks.NewProcessor(accountService)
	processor.Register(mux)

	// --- Start Scheduler in background (optional) ---
	go startScheduler(redisAddr)

	// --- Run Worker Server ---
	logger.LogInfo("🧠 Asynq Worker running on Redis: %s", redisAddr)
	if err := server.Run(mux); err != nil {
		logger.LogError("❌ Asynq worker crashed: %v", err)
	}
}

// startScheduler runs the nightly arrears calculation
func startScheduler(redisAddr string) {
	scheduler := asynq.NewScheduler(asynq.RedisClientOpt{Addr: redisAddr}, nil)

	// Run every day at 2:00 AM
	_, err := scheduler.Register("0 2 * * *",
		asynq.NewTask(tasks.TaskCalculateArrears, []byte(`{"run_type":"nightly"}`), asynq.Queue("arrears")),
	)

	if err != nil {
		log.Fatalf("❌ Failed to register scheduler: %v", err)
	}

	log.Println("📅 Nightly arrears scheduler registered (2:00 AM)...")

	if err := scheduler.Run(); err != nil {
		log.Fatalf("❌ Scheduler crashed: %v", err)
	}
}
