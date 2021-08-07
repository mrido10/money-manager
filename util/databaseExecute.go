package util

import (
	"database/sql"
	"fmt"
)

func DBInsertExsecute(query string, param ...interface{}) error {
	db, err := ConnectMySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query(query, param...)
	if err != nil {

		return err
	}
	defer rows.Close()
	return nil
}

// this funtion still not working
func DBSelectExsecute(query string, param ...interface{}) (*sql.Rows, error) {
	db, err := ConnectMySQL()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	fmt.Println(param...)

	rows, err := db.Query(query, param...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return rows, nil
}
