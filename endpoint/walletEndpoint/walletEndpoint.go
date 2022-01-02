package walletEndpoint

import (
	"money-manager/service/walletService"
	"money-manager/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type walletStruct struct {
	util.DataIn
}

var WalletController = walletStruct{}

func (validate walletStruct) GetWallet(c *gin.Context) {
	action := c.Param("action")
	switch action {
	case "getWallet":
		validate.JWTValidations(c, walletService.WalletService.GetWallet)
	case "getWalletWithPagination":
		validate.JWTValidations(c, walletService.WalletService.GetWalletWithPagination)
	case "getWalletByGroup":
		validate.JWTValidations(c, walletService.WalletService.GetListWalletByGroup)
	default:
		util.Response(c, http.StatusNotFound, "Not Found", nil)
	}
}

func (validate walletStruct) SaveWallet(c *gin.Context) {
	action := c.Param("action")
	switch action {
	case "saveWallet":
		validate.JWTValidations(c, walletService.WalletService.SaveWallet)
	default:
		util.Response(c, http.StatusNotFound, "Not Found", nil)
	}
}

func (validate walletStruct) DeleteWallet(c *gin.Context) {
	action := c.Param("action")
	switch action {
	case "deleteWallet":
		validate.JWTValidations(c, walletService.WalletService.DeleteWallet)
	default:
		util.Response(c, http.StatusNotFound, "Not Found", nil)
	}
}