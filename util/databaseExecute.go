package util

import (
	"database/sql"
	"log"
)

func DBExecute(query string, param ...interface{}) error {
	db, err := ConnectMySQL()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer db.Close()

	rows, err := db.Query(query, param...)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer rows.Close()
	return nil
}

// this funtion still not working
func DBSelectExsecute(query string, param ...interface{}) (*sql.Rows, error) {
	db, err := ConnectMySQL()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query, param...)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	return rows, nil
}
