package walletGroupService

import (
	"money-manager/dao"
	"money-manager/model"
	"money-manager/repository"

	log "github.com/sirupsen/logrus"
)

func SaveWalletGroup(c model.DataIN) (listData interface{}, responseCode int, msg string, err error) {
	var data repository.WalletGroup
	if err = c.GinContext.ShouldBindJSON(&data); err != nil {
		log.Error(err)
		responseCode, msg = 400, err.Error()
		return
	}

	walletGroupID, err := generateWalletGroupId(c.UserID)
	if err != nil {
		responseCode, msg = 400, err.Error()
		return
	}
	data.UserID = c.UserID
	data.WalletGroupID = walletGroupID

	dao.SaveWalletGroup(data)
	responseCode, msg = 200, "succes save " + data.GroupName
	return
}