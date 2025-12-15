package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init(dsn string) error {
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	return DB.Ping()
}
