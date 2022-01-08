package walletService

import (
	"log"
	"money-manager/dao"

	"money-manager/util"
)

func generateWalletId(userID string) (string, error) {
	latestWalletID, err := dao.GetLastWalletID(userID)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return util.GenerateID(latestWalletID, "WAL"), nil
}
