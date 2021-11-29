package crud

import (
	"log"
	"money-manager/dao"
	"money-manager/model/domain"
	"money-manager/util"
)

type WalletGroup interface {
	GetListWalletGroup()
	SaveWalletGroup()
}

func (d DTO) GetListWalletGroup() {
	log.Println(d)
	listData, err := dao.GetListWalletGroup(d.UserID)
	util.ToShowResp(listData, err, d.GinContext)
}

func (d DTO) SaveWalletGroup() {
	var data domain.WalletGroup
	if err := d.GinContext.ShouldBindJSON(&data); err != nil {
		log.Println(err)
		util.Response(d.GinContext, 400, err.Error(), nil)
		return
	}

	walletGroupID, err := generateWalletGroupId(d.UserID)
	if err != nil {
		util.Response(d.GinContext, 400, err.Error(), nil)
		return
	}
	data.UserID = d.UserID
	data.WalletGroupID = walletGroupID

	dao.SaveWalletGroup(data)
	util.Response(d.GinContext, 200, "succes save " + data.GroupName, nil)
}

func generateWalletGroupId(userID string) (string, error) {
	latestWalletID, err := dao.GetLastWalletGroupID(userID)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return util.GenerateID(latestWalletID, "WG"), nil
}
