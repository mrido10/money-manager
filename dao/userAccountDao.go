package dao

import (
	"money-manager/util"
)

func CheckUserAccount(userID string) (int, error) {
	db, err := util.ConnectMySQL()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	query := `SELECT COUNT(*) as total FROM userAccount WHERE userID = ?`

	rows, err := db.Query(query, userID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var total int
	for rows.Next() {
		rows.Scan(&total)
	}

	return total, nil
}

func InsertUserAccount(userId string, userName string) error {
	query := `INSERT INTO userAccount(userID, userName) VALUE(?, ?)`
	err := util.DBInsertExsecute(query, userId, userName)
	if err != nil {
		return err
	}

	return nil
}
