package dao

import (
	"fmt"
	"money-manager/repository"
	"money-manager/util"
)

type (
	CountTotalTransactions struct {
		TypeName    string `json:"typeName"`
		TotalAmount int    `json:"totalAmount"`
	}
)

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

func InsertOrUpadateTransaction(note repository.Notes) error {
	query := `REPLACE INTO notes 
			(notesID, userID, walletID, notesTypeID, notesCategoryID, notesName, 
			amount, fromAccount, toAccount, description, notesDate)
			VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	db, err := util.ConnectMySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query(query,
		note.NotesID,
		note.UserId,
		note.WalletID,
		note.NotesTypeID,
		note.NotesCategoryID,
		note.NotesName,
		note.Amount,
		note.FormAccount,
		note.ToAccount,
		note.Description,
		note.NotesDate,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
