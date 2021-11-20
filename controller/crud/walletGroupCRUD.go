package crud

import (
	"log"
	"money-manager/dao"
	"money-manager/model/domain"
	"money-manager/util"

	"github.com/gin-gonic/gin"
)

func GetListWalletGroup(c *gin.Context, userID string) {
	listData, err := dao.GetListWalletGroup(userID)
	util.ToShowResp(listData, err, c)
}

func SaveWalletGroup(c *gin.Context, userID string) {
	var data domain.WalletGroup
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Println(err)
		util.Response(c, 400, err.Error(), nil)
		return
	}

	walletGroupID, err := generateWalletGroupId(userID)
	if err != nil {
		util.Response(c, 400, err.Error(), nil)
		return
	}
	data.UserID = userID
	data.WalletGroupID = walletGroupID

	dao.SaveWalletGroup(data)
	util.Response(c, 200, "succes save "+data.GroupName, nil)
}

func generateWalletGroupId(userID string) (string, error) {
	latestWalletID, err := dao.GetLastWalletGroupID(userID)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return util.GenerateID(latestWalletID, "WG"), nil
}
