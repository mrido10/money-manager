package dao

import (
	log "github.com/sirupsen/logrus"
	"money-manager/repository"
	"money-manager/util"
)

var (
	wallet      repository.Wallet
	walletGroup repository.WalletGroup
)

func GetListWallet(userID string) ([]repository.Wallet, error) {
	query := `SELECT * FROM wallet WHERE userID = ? ORDER BY walletID DESC`

	db, err := util.ConnectMySQL()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query, userID)

	if err != nil {
		log.Error(err.Error())
		return []repository.Wallet{}, err
	}
	defer rows.Close()

	var result []repository.Wallet

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
			log.Error(err.Error())
			return []repository.Wallet{}, err
		}

		result = append(result, acc)
	}

	if err = rows.Err(); err != nil {
		log.Error(err.Error())
		return []repository.Wallet{}, err
	}

	return result, nil
}

func GetListWalletByPagin(userID string, limit int, offset int) ([]repository.Wallet, error) {
	query := `SELECT * FROM wallet WHERE userID = ? ORDER BY walletID DESC
			  LIMIT ? OFFSET ?`

	db, err := util.ConnectMySQL()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query, userID, limit, offset)

	if err != nil {
		log.Error(err.Error())
		return []repository.Wallet{}, err
	}
	defer rows.Close()

	var result []repository.Wallet

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
			log.Error(err.Error())
			return []repository.Wallet{}, err
		}

		result = append(result, acc)
	}

	if err = rows.Err(); err != nil {
		log.Error(err.Error())
		return []repository.Wallet{}, err
	}

	return result, nil
}

func GetListWalletByGroup(userID string, walletGroupID string) ([]repository.Wallet, error) {
	query := `SELECT * FROM wallet WHERE userID = ? AND walletGroupID = ? 
				ORDER BY walletGroupID ASC, walletID ASC `

	db, err := util.ConnectMySQL()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query, userID, walletGroupID)

	if err != nil {
		log.Error(err.Error())
		return []repository.Wallet{}, err
	}
	defer rows.Close()

	var result []repository.Wallet

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
			log.Error(err.Error())
			return []repository.Wallet{}, err
		}

		result = append(result, acc)
	}

	if err = rows.Err(); err != nil {
		log.Error(err.Error())
		return []repository.Wallet{}, err
	}

	return result, nil
}

func SaveWallet(wallet repository.Wallet) error {
	query := `INSERT INTO wallet(walletID, userID, walletGroupID, walletName, amount, description)
			VALUE(?, ? , ?, ?, ?, ?)`
	return util.DBExecute(query, wallet.WalletID, wallet.UserID, wallet.WalletGroupID, wallet.WalletName, wallet.Amount, wallet.Description)
}

func GetLastWalletID(userID string) (string, error) {
	query := `SELECT walletID FROM wallet ORDER BY walletID DESC LIMIT 1`

	db, err := util.ConnectMySQL()
	if err != nil {
		log.Error(err.Error())
		return "", err
	}
	defer db.Close()

	rows, err := db.Query(query)

	if err != nil {
		log.Error(err.Error())
		return "", err
	}
	defer rows.Close()

	var result string

	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			log.Error(err.Error())
			return "", err
		}
	}

	if err = rows.Err(); err != nil {
		log.Error(err.Error())
		return "", err
	}

	return result, nil
}

func DeleteWallet(walletID string, userID string) error {
	query := `DELETE FROM wallet WHERE walletID = ? AND userID = ?`
	return util.DBExecute(query, walletID, userID)
}
