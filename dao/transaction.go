package dao

import (
	"fmt"
	"money-manager/util"
)

type CountTotalTransactions struct {
	TypeName    string `json:"typeName"`
	TotalAmount int    `json:"totalAmount"`
}

func GetCountTotalTransactions(userID string) (map[string]CountTotalTransactions, error) {
	query := `SELECT tp.typeName, SUM(tr.amount) as totalAmount FROM transaction tr
			INNER JOIN transactiontype tp
			ON tp.typeID = tr.transactionTypeID
			WHERE tr.userID = ?
			GROUP BY tr.transactionTypeID`
	db, err := util.ConnectMySQL()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query, userID)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]CountTotalTransactions)
	for rows.Next() {
		trans := CountTotalTransactions{}
		err = rows.Scan(
			&trans.TypeName,
			&trans.TotalAmount,
		)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		result[trans.TypeName] = trans
	}

	return result, nil
}
