package domain

import (
	"time"
)

type Transaction struct {
	TransactionID         int       `json:"transactionID"`
	TransactionTypeID     int       `json:"transactionTypeID"`
	TransactionCategoryID int       `json:"transactionCategoryID"`
	AccountID             string    `json:"accountID"`
	Amount                int64     `json:"amount"`
	FormAccount           int       `json:"fromAccount"`
	ToAccount             int       `json:"toAccount"`
	Description           string    `json:"description"`
	TransactionDate       time.Time `json:"transactionData"`
}
