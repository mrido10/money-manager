package controller

import (
	"fmt"
	"money-manager/dao"
	"money-manager/util"

	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context) {
	listName := c.Query("listName")

	switch listName {
	case "account":
		getListAccount(c)
	case "accountGroup":
		getListAccountGroup(c)
	}
}

func getListAccount(c *gin.Context) {
	listData, err := dao.GetListAccount()
	toShowResp(listData, err, c)
}

func getListAccountGroup(c *gin.Context) {
	listData, err := dao.GetListAccountGroup()
	toShowResp(listData, err, c)
}

func toShowResp(data interface{}, err error, c *gin.Context) {
	if err != nil {
		fmt.Println(err.Error())
		util.Response(c, 400, err.Error(), nil)
		return
	}

	util.Response(c, 200, "succes", data)
}
