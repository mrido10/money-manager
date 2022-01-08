package controller

import (
	"fmt"
	"money-manager/config"
	"money-manager/endpoint/walletEndpoint"
	"money-manager/endpoint/walletGroupEndpoint"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Route(route *gin.Engine) {
	walletRoute := "/wallet/:action"
	route.GET(walletRoute, walletEndpoint.WalletEndpoint.GetWallet)
	route.POST(walletRoute, walletEndpoint.WalletEndpoint.SaveWallet)
	route.DELETE(walletRoute, walletEndpoint.WalletEndpoint.DeleteWallet)

	walletGroupRoute := "/walletGroup/:action"
	route.GET(walletGroupRoute, walletGroupEndpoint.WalletGroupEndpoint.GetWalletGroup)
	route.POST(walletGroupRoute, walletGroupEndpoint.WalletGroupEndpoint.SaveWalletGroup)

	creditPrint()
	conf, err := config.GetConfig()
	if err != nil {
		log.Error(err)
		return 
	}

	if err := route.Run(":" + conf.Server.Port); err != nil {
		panic(err)
	}
}

func creditPrint(){
	fmt.Print("\n\n")
	fmt.Println("\t@@@                   @@@     @@@                @@@@      @@@         ")
	fmt.Println("\t @@                    @@@   @@@                  @@       @@          ")
	fmt.Println("\t @@                    @@@@ @@@@                  @@       @@          ")
	fmt.Println("\t @@@@    @@@     @@    @@ @@@ @@   @@@ @@@        @@     @@@@     @@@  ")
	fmt.Println("\t @@@@@@    @@   @@     @@  @  @@    @@@   @       @@   @@@@@@   @@@@@@@")
	fmt.Println("\t @@  @@     @@ @@      @@     @@    @@            @@   @@  @@   @@   @@")
	fmt.Println("\t @@@@@@      @@@       @@     @@    @@       @@   @@   @@@@@@   @@@@@@@")
	fmt.Println("\t@@@@@        @@       @@@     @@@  @@@       @@  @@@@    @@@@@    @@@  ")
	fmt.Println("\t            @@")
	fmt.Print("\t          @@@\n\n")
}
