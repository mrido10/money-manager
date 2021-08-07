package view

import (
	"money-manager/controller"
	"github.com/gin-gonic/gin"
)

func StartService() {
	route := gin.Default()

	transactionRoute(route)
	registerRoute(route)
	getList(route)

	if err := route.Run(":3002"); err != nil {
		panic(err)
	}
}

func registerRoute(route *gin.Engine) {
	route.POST("/money-manager/register", controller.RegisterTo)
}

func transactionRoute(route *gin.Engine) {
	route.GET("/money-manager/getTransactionNote", controller.GetAllTrancsactionNote)
}

func getList(route *gin.Engine) {
	route.GET("/money-manager/getList", controller.GetList)
}
