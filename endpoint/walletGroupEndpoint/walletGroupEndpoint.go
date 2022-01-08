package walletGroupEndpoint

import (
	"money-manager/service/walletGroupService"
	"money-manager/util"

	"github.com/gin-gonic/gin"
)

type walletGroupStruct struct {
	util.DataIn
}

var WalletGroupEndpoint = walletGroupStruct{}

func (validate walletGroupStruct) GetWalletGroup(c *gin.Context) {
	action := c.Param("action")
	switch action {
	case "getWalletGroup":
		validate.JWTValidations(c, walletGroupService.GetListWalletGroup)
	default:
		util.ResponseNotFound(c)
	}
}

func (validate walletGroupStruct) SaveWalletGroup(c *gin.Context) {
	action := c.Param("action")
	switch action {
	case "saveWalletGroup":
		validate.JWTValidations(c, walletGroupService.SaveWalletGroup)
	default:
		util.ResponseNotFound(c)
	}
}