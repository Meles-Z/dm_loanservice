package main

import (
	"context"
	cmdGrpc "dm_loanservice/cmd/grpc"
	"dm_loanservice/cmd/rest"
	"dm_loanservice/drivers/dbmigrations"
	"dm_loanservice/drivers/goconf"
	"dm_loanservice/drivers/jwt"
	"dm_loanservice/drivers/logger"
	"dm_loanservice/drivers/postgres"
	redisLib "dm_loanservice/drivers/redis"
	"dm_loanservice/drivers/validator"
	"dm_loanservice/internal/endpoint"
	accountR "dm_loanservice/internal/service/repository/account"
	accountAuditLogR "dm_loanservice/internal/service/repository/account_audit_log"
	accountflagR "dm_loanservice/internal/service/repository/account_flag"
	accountLockRuleR "dm_loanservice/internal/service/repository/account_lock_rule"
	collateralR "dm_loanservice/internal/service/repository/collateral"
	duediligenceR "dm_loanservice/internal/service/repository/due_diligence"
	investorRestrictionR "dm_loanservice/internal/service/repository/investor_restriction"
	lateFeeRuleR "dm_loanservice/internal/service/repository/late_fee_rule"
	securitisationR "dm_loanservice/internal/service/repository/securitisation"
	serviceRestrictionR "dm_loanservice/internal/service/repository/service_restriction"
	accountSvc "dm_loanservice/internal/service/usecase/account"
	accountAuditLogSvc "dm_loanservice/internal/service/usecase/account_audit_log"
	accountflagSvc "dm_loanservice/internal/service/usecase/account_flag"
	accountLockRuleSvc "dm_loanservice/internal/service/usecase/account_lock_rule"
	duediligenceSvc "dm_loanservice/internal/service/usecase/due_diligence"
	investorRestrictionSvc "dm_loanservice/internal/service/usecase/investor_restriction"
	lateFeeRuleSvc "dm_loanservice/internal/service/usecase/late_fee_rule"
	securitisationSvc "dm_loanservice/internal/service/usecase/securitisation"
	serviceRestrictionSvc "dm_loanservice/internal/service/usecase/service_restriction"

	"fmt"

	_ "github.com/lib/pq"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
	"github.com/pressly/goose"
)

// @title DMortgages - User Service
// @version 1.0.0
// @description DMortgages - User Service
func main() {
	var (
		err error
		ctx = context.Background()
	)

	goconf.Config() // init config

	// initialize logger
	logger.InitLogger()

	// initialize validator
	validator.InitValidator()

	// DB connection
	pgdb, err := postgres.NewDBMaster()
	if err != nil {
		panic(fmt.Errorf("error connect postgres: %v", err))
	}
	defer func() {
		_ = pgdb.Close()
	}()
	// run DB migrations
	dbGoose, err := dbmigrations.RunDBMigrations()
	if err != nil {
		panic(fmt.Errorf("error run DB migrations: %v", err))
	}
	// Set dialect for goose
	if err := goose.SetDialect("postgres"); err != nil {
		panic(fmt.Errorf("error setting goose dialect: %v", err))
	}
	// Get current version (this will create table if needed without panic)
	currentVersion, err := goose.GetDBVersion(dbGoose)
	if err != nil {
		// If table doesn't exist, GetDBVersion will create it
		fmt.Printf("Migration table initialization: %v\n", err)
	}
	fmt.Printf("Current DB migration version: %d\n", currentVersion)

	// Run pending migrations
	if err := goose.Up(dbGoose, goconf.Config().GetString("migrations_folder")); err != nil {
		panic(fmt.Errorf("error running migrations: %v", err))
	}

	// setup customer service wrapper (gRPC client)
	// set expiry jwt
	jwt.SetExpiry(goconf.Config().GetInt("jwt.expire"))

	// initialize repository
	lateFeeRuleRepo := lateFeeRuleR.NewRepository(pgdb)
	accountRepo := accountR.NewRepository(pgdb)
	duediligenceRepo := duediligenceR.NewRepository(pgdb)
	accountFlagRepo := accountflagR.NewRepository(pgdb)
	accountAuditLogRepo := accountAuditLogR.NewRepository(pgdb)
	serviceRestrictionRepo := serviceRestrictionR.NewRepository(pgdb)
	investorRestrictionRepo := investorRestrictionR.NewRepository(pgdb)
	accountLockRuleRepo := accountLockRuleR.NewRepository(pgdb)
	collateralRepo := collateralR.NewRepository(pgdb)
	securitisationRepo := securitisationR.NewRepository(pgdb)
	redisConn := redisLib.GetConnection(context.Background())
	_ = redisConn

	// initialize endpoints
	e := endpoint.NewEndpoints(
		lateFeeRuleSvc.NewService(lateFeeRuleRepo),
		accountSvc.NewService(accountRepo, duediligenceRepo, accountLockRuleRepo, serviceRestrictionRepo),
		duediligenceSvc.NewService(duediligenceRepo, accountRepo, accountFlagRepo, accountAuditLogRepo),
		accountflagSvc.NewService(accountFlagRepo, accountRepo, accountAuditLogRepo),
		accountAuditLogSvc.NewService(accountAuditLogRepo, accountRepo),
		accountLockRuleSvc.NewService(accountLockRuleRepo, accountRepo),
		serviceRestrictionSvc.NewService(serviceRestrictionRepo, accountRepo, investorRestrictionRepo, accountLockRuleRepo),
		investorRestrictionSvc.NewService(investorRestrictionRepo, accountRepo),
		securitisationSvc.NewService(securitisationRepo, accountRepo, nil, duediligenceRepo, accountFlagRepo, collateralRepo),
	)

	// define gothic provider
	goth.UseProviders(
		google.New(
			goconf.Config().GetString("provider.google.key"),
			goconf.Config().GetString("provider.google.secret"),
			goconf.Config().GetString("provider.google.callbackURL"),
			goconf.Config().GetStringSlice("provider.google.scope")...,
		),
		facebook.New(
			goconf.Config().GetString("provider.facebook.app_id"),
			goconf.Config().GetString("provider.facebook.secret"),
			goconf.Config().GetString("provider.facebook.callbackURL"),
			goconf.Config().GetStringSlice("provider.facebook.scope")...,
		),
	)

	// run rest server
	go func() {
		_ = rest.RunServer(ctx, e, goconf.Config().GetString("rest.port"))
	}()

	// run grpc server
	grpcServer := cmdGrpc.RunServer(
		ctx,
		goconf.Config().GetString("grpc.port"),
	)
	if grpcServer != nil {
		panic(fmt.Errorf("error starting gRPC server: %v", grpcServer))
	}

}
