package repository

type Wallet struct {
	WalletID      string `json:"walletID"`
	UserID        string `json:"userID"`
	WalletGroupID string `json:"walletGroupID"`
	WalletName    string `json:"walletName" default:""`
	Amount        int64  `json:"amount" default:""`
	Description   string `json:"description" default:""`
}
