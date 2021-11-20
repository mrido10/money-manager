package crud

import (
	"log"
	"money-manager/dao"
	"money-manager/model/domain"
	"money-manager/util"

	"github.com/gin-gonic/gin"
)

func GetListWallet(c *gin.Context, userID string) {
	listData, err := dao.GetListWallet(userID)
	util.ToShowResp(listData, err, c)
}

func GetListWalletByGroup(c *gin.Context, userID string) {
	var data struct {
		WalletGroupID string `json:"walletGroupID"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Println(err)
		util.Response(c, 400, err.Error(), nil)
		return
	}
	listData, err := dao.GetListWalletByGroup(userID, data.WalletGroupID)
	util.ToShowResp(listData, err, c)
}

func SaveWallet(c *gin.Context, userID string) {
	var data domain.Wallet
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Println(err)
		util.Response(c, 400, err.Error(), nil)
		return
	}

	walletID, err := generateWalletId(userID)
	if err != nil {
		util.Response(c, 400, err.Error(), nil)
		return
	}
	data.UserID = userID
	data.WalletID = walletID

	dao.SaveWallet(data)
	util.Response(c, 200, "succes save "+data.WalletName, nil)
}

func generateWalletId(userID string) (string, error) {
	latestWalletID, err := dao.GetLastWalletID(userID)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return util.GenerateID(latestWalletID, "WAL"), nil
}
