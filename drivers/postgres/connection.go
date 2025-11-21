package postgres

import (
	"fmt"

	"dm_loanservice/drivers/goconf"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDBMaster() (*sqlx.DB, error) {
	var (
		host     = goconf.Config().GetString("postgres.master.host")
		port     = goconf.Config().GetString("postgres.master.port")
		user     = goconf.Config().GetString("postgres.master.user")
		password = goconf.Config().GetString("postgres.master.password")
		dbname   = goconf.Config().GetString("postgres.master.db")
	)
	DBMS := goconf.Config().GetString("postgres.master.driver")
	psqlMaster := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Connect(DBMS, psqlMaster)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewDBSlave() (*sqlx.DB, error) {
	var (
		host     = goconf.Config().GetString("postgres.slave.host")
		port     = goconf.Config().GetString("postgres.slave.port")
		user     = goconf.Config().GetString("postgres.slave.user")
		password = goconf.Config().GetString("postgres.slave.password")
		dbname   = goconf.Config().GetString("postgres.slave.db")
	)

	DBMS := goconf.Config().GetString("postgres.slave.driver")
	psqlSlave := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Connect(DBMS, psqlSlave)
	if err != nil {
		return nil, err
	}

	return db, nil
}
