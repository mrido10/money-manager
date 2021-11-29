package dao

import (
	"money-manager/util"
	"strings"
	"time"
)

func RegisterToAccountGroup(name string) error {
	query := `INSERT INTO walletGroup (groupName)
			VALUE(?)`
	return util.DBExecute(query, name)
}

func RegisterToAccount(groupID int, accName string, amount int, desc string) error {
	t := time.Now()
	tString := t.Format(("2006-01-02 15:04:05.000"))
	tString = strings.Replace(tString, "-", "", 2)
	tString = strings.Replace(tString, " ", "", 1)
	tString = strings.Replace(tString, ":", "", 2)
	tString = strings.Replace(tString, ".", "", 1)
	tString = tString[2 : len(tString)-1]

	query := `INSERT INTO account (accountID, groupID, accountName, amount, description)
			VALUE(?, ?, ?, ?, ?)`

	var accNameSub *string = &accName
	id := accName[:3] + tString
	return util.DBExecute(query, id, groupID, *accNameSub, amount, desc)
}

func RegisterToTransactionType(name string) error {
	query := `INSERT INTO transactionType (typeName)
			VALUE(?)`
	return util.DBExecute(query, name)
}

func RegisterToTransactionCategory(name string) error {
	query := `INSERT INTO transactionCategory (categoryName)
	VALUE(?)`
	return util.DBExecute(query, name)
}
