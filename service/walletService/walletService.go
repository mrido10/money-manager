package walletService

import (
	"money-manager/dao"

	log "github.com/sirupsen/logrus"

	"money-manager/util"
)

func generateWalletId(userID string) (string, error) {
	latestWalletID, err := dao.GetLastWalletID(userID)
	if err != nil {
		log.Error(err)
		return "", err
	}

	return util.GenerateID(latestWalletID, "WAL"), nil
}
