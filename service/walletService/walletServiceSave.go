package walletService

import (
	"log"
	"money-manager/dao"
	"money-manager/model"
	"money-manager/model/domain"
)

func (d data) SaveWallet(c model.DataIN) (listData interface{}, responseCode int, msg string, err error) {
	var dtWallet domain.Wallet
	if err = c.GinContext.ShouldBindJSON(&dtWallet); err != nil {
		log.Println(err)
		return
	}

	walletID, errr := generateWalletId(c.UserID)
	if errr != nil {
		log.Println(errr)
		msg = errr.Error()
		err = errr
		return
	}
	dtWallet.UserID = c.UserID
	dtWallet.WalletID = walletID

	err = dao.SaveWallet(dtWallet)
	if err != nil {
		msg = err.Error()
		return
	}
	responseCode = 200
	msg = "succes save " + dtWallet.WalletName
	return
}