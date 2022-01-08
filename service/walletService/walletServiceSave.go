package walletService

import (
	"money-manager/dao"
	"money-manager/model"
	"money-manager/repository"

	log "github.com/sirupsen/logrus"
)

func SaveWallet(c model.DataIN) (listData interface{}, responseCode int, msg string, err error) {
	var dtWallet repository.Wallet
	if err = c.GinContext.ShouldBindJSON(&dtWallet); err != nil {
		log.Error(err)
		return
	}

	walletID, errr := generateWalletId(c.UserID)
	if errr != nil {
		log.Error(errr)
		responseCode, msg, err = 400, errr.Error(), errr
		return
	}
	dtWallet.UserID = c.UserID
	dtWallet.WalletID = walletID

	err = dao.SaveWallet(dtWallet)
	if err != nil {
		responseCode, msg= 400, err.Error()
		return
	}
	responseCode, msg = 200, "succes save " + dtWallet.WalletName
	return
}