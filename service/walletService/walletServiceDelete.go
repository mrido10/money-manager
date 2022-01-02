package walletService

import (
	"log"
	"money-manager/dao"
	"money-manager/model"
)

func (d data) DeleteWallet(c model.DataIN) (listData interface{}, responseCode int, msg string, err error) {
	var dtWallet struct {
		WalletID string `json:"walletID"`
	}
	if err = c.GinContext.ShouldBindJSON(&dtWallet); err != nil {
		log.Println(err)
		return
	}
	err = dao.DeleteWallet(dtWallet.WalletID, c.UserID)
	if err != nil {
		log.Println(err)
		return
	}
	responseCode = 200
	msg = "succes delete " + dtWallet.WalletID
	return
}