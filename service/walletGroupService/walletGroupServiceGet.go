package walletGroupService

import (
	"money-manager/dao"
	"money-manager/model"
)

func GetListWalletGroup(c model.DataIN) (listData interface{}, responseCode int, msg string, err error) {
	listData, err = dao.GetListWalletGroup(c.UserID)
	responseCode = 200
	return 
}