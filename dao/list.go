package dao

import (
	"fmt"
	"money-manager/model/domain"
	"money-manager/util"
)

var (
	account      domain.Account
	accountGroup domain.AccountGroup
)

func GetListAccount() ([]domain.Account, error) {
	query := `SELECT * FROM account`

	db, err := util.ConnectMySQL()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err.Error())
		return []domain.Account{}, err
	}
	defer rows.Close()

	var result []domain.Account

	for rows.Next() {
		acc := account
		err = rows.Scan(
			&acc.AccountID,
			&acc.GroupID,
			&acc.AccountName,
			&acc.Amount,
			&acc.Description,
		)

		if err != nil {
			fmt.Println(err.Error())
			return []domain.Account{}, err
		}

		result = append(result, acc)
	}

	if err = rows.Err(); err != nil {
		return []domain.Account{}, err
	}

	return result, nil
}

func GetListAccountGroup() ([]domain.AccountGroup, error) {
	query := `SELECT * FROM accountGroup`

	db, err := util.ConnectMySQL()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if err != nil {
		fmt.Println(err.Error())
		return []domain.AccountGroup{}, err
	}

	var result []domain.AccountGroup

	for rows.Next() {
		acc := accountGroup
		err = rows.Scan(
			&acc.GroupID,
			&acc.GroupName,
		)

		if err != nil {
			fmt.Println(err.Error())
			return []domain.AccountGroup{}, err
		}

		result = append(result, acc)
	}

	if err = rows.Err(); err != nil {
		return []domain.AccountGroup{}, err
	}

	return result, nil
}
