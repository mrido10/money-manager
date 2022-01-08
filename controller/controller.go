package controller

import (
	"log"
	"money-manager/config"
	"money-manager/endpoint/walletEndpoint"
	"money-manager/endpoint/walletGroupEndpoint"

	"github.com/gin-gonic/gin"
)

func Route(route *gin.Engine) {
	walletRoute := "/wallet/:action"
	route.GET(walletRoute, walletEndpoint.WalletEndpoint.GetWallet)
	route.POST(walletRoute, walletEndpoint.WalletEndpoint.SaveWallet)
	route.DELETE(walletRoute, walletEndpoint.WalletEndpoint.DeleteWallet)

	walletGroupRoute := "/walletGroup/:action"
	route.GET(walletGroupRoute, walletGroupEndpoint.WalletGroupEndpoint.GetWalletGroup)
	route.POST(walletGroupRoute, walletGroupEndpoint.WalletGroupEndpoint.SaveWalletGroup)

	conf, err := config.GetConfig()
	if err != nil {
		log.Println(err)
		return 
	}

	if err := route.Run(":" + conf.Server.Port); err != nil {
		panic(err)
	}
}
