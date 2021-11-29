package util

import (
	"database/sql"
	"log"

	"money-manager/config"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL() (*sql.DB, error) {
	c, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db, err := sql.Open("mysql", c.MySql.DataSource)
	if err != nil {
		return nil, err
	}

	return db, nil
}
