package walletEndpoint

import (
	"money-manager/service/walletService"
	"money-manager/util"

	"github.com/gin-gonic/gin"
)

type walletStruct struct {
	util.DataIn
}

var WalletEndpoint = walletStruct{}

func (validate walletStruct) GetWallet(c *gin.Context) {
	action := c.Param("action")
	switch action {
	case "getWallet":
		validate.JWTValidations(c, walletService.GetWallet)
	case "getWalletWithPagination":
		validate.JWTValidations(c, walletService.GetWalletWithPagination)
	case "getWalletByGroup":
		validate.JWTValidations(c, walletService.GetListWalletByGroup)
	default:
		util.ResponseNotFound(c)
	}
}

func (validate walletStruct) SaveWallet(c *gin.Context) {
	action := c.Param("action")
	switch action {
	case "saveWallet":
		validate.JWTValidations(c, walletService.SaveWallet)
	default:
		util.ResponseNotFound(c)
	}
}

func (validate walletStruct) DeleteWallet(c *gin.Context) {
	action := c.Param("action")
	switch action {
	case "deleteWallet":
		validate.JWTValidations(c, walletService.DeleteWallet)
	default:
		util.ResponseNotFound(c)
	}
}