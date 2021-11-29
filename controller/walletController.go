package controller

import (
	"log"
	"money-manager/controller/crud"
	"money-manager/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	wallet      crud.Wallet
	walletGroup crud.WalletGroup
)
func GetList(c *gin.Context) {
	claims, err := util.Authorization(c)
	if err != nil {
		log.Println(err.Error())
		util.Response(c, 401, "Unauthorized", nil)
		return
	}

	userID := claims["id"].(string)
	action := c.Param("action")

	wallet = crud.DTO{
		GinContext: c, 
		UserID: userID,
	}

	walletGroup = crud.DTO{
		GinContext: c, 
		UserID: userID,
	}

	switch action {
		case "getWallet":
			wallet.GetListWallet()
		case "getWalletByGroup":
			wallet.GetListWalletByGroup()
		case "getWalletGroup":
			walletGroup.GetListWalletGroup()
		default:
			util.Response(c, http.StatusNotFound, "Not Found", nil)
	}
}

func SaveWallet(c *gin.Context) {
	claims, err := util.Authorization(c)
	if err != nil {
		log.Println(err.Error())
		util.Response(c, 401, "Unauthorized", nil)
		return
	}

	userID := claims["id"].(string)
	action := c.Param("action")
	
	wallet = crud.DTO{
		GinContext: c, 
		UserID: userID,
	}

	walletGroup = crud.DTO{
		GinContext: c, 
		UserID: userID,
	}

	switch action {
		case "saveWallet":
			wallet.SaveWallet()
		case "saveWalletGroup":
			walletGroup.SaveWalletGroup()
		case "deleteWallet":
			wallet.DeleteWallet()
		default:
			util.Response(c, http.StatusNotFound, "Not Found", nil)
	}
}
