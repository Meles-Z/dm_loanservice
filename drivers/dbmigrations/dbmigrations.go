package dbmigrations

import (
	"database/sql"
	"dm_loanservice/drivers/goconf"
	"fmt"

	"github.com/pressly/goose"
)

func RunDBMigrations() (*sql.DB, error) {
	var (
		host     = goconf.Config().GetString("postgres.master.host")
		port     = goconf.Config().GetString("postgres.master.port")
		user     = goconf.Config().GetString("postgres.master.user")
		password = goconf.Config().GetString("postgres.master.password")
		dbname   = goconf.Config().GetString("postgres.master.db")
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dbGoose, err := goose.OpenDBWithDriver("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return dbGoose, nil
}
