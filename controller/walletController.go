package controller

import (
	"log"
	"money-manager/controller/crud"
	"money-manager/util"

	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context) {
	claims, err := util.Authorization(c)
	if err != nil {
		log.Println(err.Error())
		util.Response(c, 401, "Unauthorized", nil)
		return
	}

	userID := claims["id"].(string)
	paramName := c.Query("name")

	switch paramName {
	case "getWallet":
		crud.GetListWallet(c, userID)
	case "getWalletByGroup":
		crud.GetListWalletByGroup(c, userID)
	case "getWalletGroup":
		crud.GetListWalletGroup(c, userID)
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
	paramName := c.Query("name")
	switch paramName {
	case "saveWallet":
		crud.SaveWallet(c, userID)
	case "saveWalletGroup":
		crud.SaveWalletGroup(c, userID)
	}
}
