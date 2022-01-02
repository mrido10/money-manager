package controller

import (
	"log"
	"money-manager/config"
	"money-manager/endpoint/walletEndpoint"

	"github.com/gin-gonic/gin"
)

func Route(route *gin.Engine) {
	route.GET("/wallet/:action", walletEndpoint.WalletController.GetWallet)
	route.POST("/wallet/:action", walletEndpoint.WalletController.SaveWallet)
	route.DELETE("/wallet/:action", walletEndpoint.WalletController.DeleteWallet)

	conf, err := config.GetConfig()
	if err != nil {
		log.Println(err)
		return 
	}

	log.Println(conf.Server.Port)
	if err := route.Run(":" + conf.Server.Port); err != nil {
		panic(err)
	}
}
