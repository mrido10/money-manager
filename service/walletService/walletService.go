package walletService

import (
	"log"
	"money-manager/dao"

	// "money-manager/model"
	"money-manager/util"
)

type data struct {
	// model.DataIN
}

var WalletService = data{}

func generateWalletId(userID string) (string, error) {
	latestWalletID, err := dao.GetLastWalletID(userID)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return util.GenerateID(latestWalletID, "WAL"), nil
}
