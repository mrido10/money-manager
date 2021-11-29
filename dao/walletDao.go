package dao

import (
	"log"
	"money-manager/model/domain"
	"money-manager/util"
)

var (
	wallet      domain.Wallet
	walletGroup domain.WalletGroup
)

func GetListWallet(userID string) ([]domain.Wallet, error) {
	query := `SELECT * FROM wallet WHERE userID = ? ORDER BY walletID DESC`

	db, err := util.ConnectMySQL()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query, userID)

	if err != nil {
		log.Println(err.Error())
		return []domain.Wallet{}, err
	}
	defer rows.Close()

	var result []domain.Wallet

	for rows.Next() {
		acc := wallet
		err = rows.Scan(
			&acc.WalletID,
			&acc.UserID,
			&acc.WalletGroupID,
			&acc.WalletName,
			&acc.Amount,
			&acc.Description,
		)

		if err != nil {
			log.Println(err.Error())
			return []domain.Wallet{}, err
		}

		result = append(result, acc)
	}

	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		return []domain.Wallet{}, err
	}

	return result, nil
}

func GetListWalletByGroup(userID string, walletGroupID string) ([]domain.Wallet, error) {
	query := `SELECT * FROM wallet WHERE userID = ? AND walletGroupID = ? 
				ORDER BY walletGroupID ASC, walletID ASC `

	db, err := util.ConnectMySQL()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query, userID, walletGroupID)

	if err != nil {
		log.Println(err.Error())
		return []domain.Wallet{}, err
	}
	defer rows.Close()

	var result []domain.Wallet

	for rows.Next() {
		acc := wallet
		err = rows.Scan(
			&acc.WalletID,
			&acc.UserID,
			&acc.WalletGroupID,
			&acc.WalletName,
			&acc.Amount,
			&acc.Description,
		)

		if err != nil {
			log.Println(err.Error())
			return []domain.Wallet{}, err
		}

		result = append(result, acc)
	}

	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		return []domain.Wallet{}, err
	}

	return result, nil
}

func SaveWallet(wallet domain.Wallet) error {
	query := `INSERT INTO wallet(walletID, userID, walletGroupID, walletName, amount, description)
			VALUE(?, ? , ?, ?, ?, ?)`
	return util.DBExecute(query, wallet.WalletID, wallet.UserID, wallet.WalletGroupID, wallet.WalletName, wallet.Amount, wallet.Description)
}

func GetLastWalletID(userID string) (string, error) {
	query := `SELECT walletID FROM wallet WHERE userID = ? ORDER BY walletID DESC LIMIT 1`

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

func DeleteWallet(walletID string, userID string) error {
	query := `DELETE FROM wallet WHERE walletID = ? AND userID = ?`
	return util.DBExecute(query, walletID, userID)
}
