package walletGroupService

import (
	"money-manager/dao"
	"money-manager/util"

	log "github.com/sirupsen/logrus"
)

func generateWalletGroupId(userID string) (string, error) {
	latestWalletID, err := dao.GetLastWalletGroupID(userID)
	if err != nil {
		log.Error(err)
		return "", err
	}

	return util.GenerateID(latestWalletID, "WG"), nil
}