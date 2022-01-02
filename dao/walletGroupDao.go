package dao

import (
	"log"
	"money-manager/repository"
	"money-manager/util"
)

func GetListWalletGroup(userID string) ([]repository.WalletGroup, error) {
	query := `SELECT * FROM walletgroup WHERE userID = ?`

	db, err := util.ConnectMySQL()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if err != nil {
		log.Println(err.Error())
		return []repository.WalletGroup{}, err
	}

	var result []repository.WalletGroup

	for rows.Next() {
		acc := walletGroup
		err = rows.Scan(
			&acc.WalletGroupID,
			&acc.UserID,
			&acc.GroupName,
		)

		if err != nil {
			log.Println(err.Error())
			return []repository.WalletGroup{}, err
		}

		result = append(result, acc)
	}

	if err = rows.Err(); err != nil {
		return []repository.WalletGroup{}, err
	}

	return result, nil
}

func GetLastWalletGroupID(userID string) (string, error) {
	query := `SELECT walletGroupID FROM walletgroup WHERE userID = ? ORDER BY walletGroupID DESC LIMIT 1`

	db, err := util.ConnectMySQL()
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	defer db.Close()

	rows, err := db.Query(query, userID)

	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	defer rows.Close()

	var result string

	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			log.Println(err.Error())
			return "", err
		}
	}

	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		return "", err
	}

	return result, nil
}

func SaveWalletGroup(wallet repository.WalletGroup) error {
	query := `INSERT INTO walletgroup(walletGroupID, userID, groupName)
			VALUE(?, ? , ?)`
	return util.DBExecute(query, wallet.WalletGroupID, wallet.UserID, wallet.GroupName)
}
