package domain

type Account struct {
	AccountID   string `json:"accountID"`
	GroupID     int    `json:"groupID"`
	AccountName string `json:"accountName"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
}
