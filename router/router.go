package view

import (
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
	route.POST("/transaction", controller.AddTransaction)

	route.GET("/getList", controller.GetList)
	route.GET("/getTotalTransactions", controller.GetListTotalCountTransaction)
	route.GET("/verify", controller.Verify)

	if err := route.Run(":3002"); err != nil {
		panic(err)
	}
}
