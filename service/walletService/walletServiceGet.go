package walletService

import (
	"money-manager/dao"
	"money-manager/model"
	"money-manager/util"
)

func (d data) GetWallet(c model.DataIN) (listData interface{}, responseCode int, msg string, err error) {
	listData, err = dao.GetListWallet(c.UserID)
	responseCode = 200
	return 
}

func (d data) GetWalletWithPagination(c model.DataIN) (listData interface{}, responseCode int, msg string, err error) {
	isPaginateExist, msg, limit, offset := util.IsExistParameterLimitAndOffset(c.GinContext)
	if !isPaginateExist {
		return
	}

	listData, err = dao.GetListWalletByPagin(c.UserID, limit, offset)
	responseCode = 200
	return 
}

func (d data) GetListWalletByGroup(c model.DataIN) (listData interface{}, responseCode int, msg string, err error) {
	var dtByGroup struct {
		WalletGroupID string `json:"walletGroupID"`
	}

	dtByGroup.WalletGroupID = c.GinContext.Query("walletGroupID")
	if dtByGroup.WalletGroupID == "" {
		responseCode = 400
		msg = "Parameter walletGroupID kosong"
		return
	}

	listData, err = dao.GetListWalletByGroup(c.UserID, dtByGroup.WalletGroupID)
	if err == nil {
		responseCode = 200
	}
	
	return
}