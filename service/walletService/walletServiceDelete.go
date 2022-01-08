package walletService

import (
	"money-manager/dao"
	"money-manager/model"

	log "github.com/sirupsen/logrus"
)

func DeleteWallet(c model.DataIN) (listData interface{}, responseCode int, msg string, err error) {
	var dtWallet struct {
		WalletID string `json:"walletID"`
	}
	if err = c.GinContext.ShouldBindJSON(&dtWallet); err != nil {
		log.Error(err)
		return
	}
	err = dao.DeleteWallet(dtWallet.WalletID, c.UserID)
	if err != nil {
		log.Error(err)
		return
	}
	responseCode = 200
	msg = "succes delete " + dtWallet.WalletID
	return
}