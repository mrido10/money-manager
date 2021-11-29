package crud

import (
	"log"
	"money-manager/dao"
	"money-manager/model/domain"
	"money-manager/util"
)

type Wallet interface {
	GetListWallet()
	GetListWalletByGroup()
	SaveWallet()
}


func (d DTO) GetListWallet() {
	listData, err := dao.GetListWallet(d.UserID)
	util.ToShowResp(listData, err, d.GinContext)
}

func (d DTO) GetListWalletByGroup() {
	var data struct {
		WalletGroupID string `json:"walletGroupID"`
	}
	if err := d.GinContext.ShouldBindJSON(&data); err != nil {
		log.Println(err)
		util.Response(d.GinContext, 400, err.Error(), nil)
		return
	}
	listData, err := dao.GetListWalletByGroup(d.UserID, data.WalletGroupID)
	util.ToShowResp(listData, err, d.GinContext)
}

func (d DTO) SaveWallet() {
	var data domain.Wallet
	if err := d.GinContext.ShouldBindJSON(&data); err != nil {
		log.Println(err)
		util.Response(d.GinContext, 400, err.Error(), nil)
		return
	}

	walletID, err := generateWalletId(d.UserID)
	if err != nil {
		util.Response(d.GinContext, 400, err.Error(), nil)
		return
	}
	data.UserID = d.UserID
	data.WalletID = walletID

	dao.SaveWallet(data)
	util.Response(d.GinContext, 200, "succes save " + data.WalletName, nil)
}

func generateWalletId(userID string) (string, error) {
	latestWalletID, err := dao.GetLastWalletID(userID)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return util.GenerateID(latestWalletID, "WAL"), nil
}
