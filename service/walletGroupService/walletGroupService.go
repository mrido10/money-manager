package walletGroupService

import (
	"log"
	"money-manager/dao"
	"money-manager/util"
)

func generateWalletGroupId(userID string) (string, error) {
	latestWalletID, err := dao.GetLastWalletGroupID(userID)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return util.GenerateID(latestWalletID, "WG"), nil
}