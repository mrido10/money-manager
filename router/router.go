package view

import (
	"log"
	"money-manager/config"
	"money-manager/controller"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartService() {
	route := gin.Default()
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	route.POST("/register", controller.RegisterTo)
	route.POST("/registerUserAccount", controller.SaveUserAccount)
	route.POST("/notes", controller.AddOrUpdateNotes)
	route.POST("/wallet", controller.SaveWallet)

	route.GET("/wallet", controller.GetList)
	route.GET("/getTotalTransactions", controller.GetListTotalCountTransaction)
	route.GET("/verify", controller.Verify)

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
