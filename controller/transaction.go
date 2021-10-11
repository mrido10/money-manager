package controller

import (
	"fmt"
	"money-manager/dao"
	"money-manager/util"

	"github.com/gin-gonic/gin"
)

type totalCountTransactions struct {
	Income  dao.CountTotalTransactions `json:"income"`
	Expence dao.CountTotalTransactions `json:"expence"`
	Total   dao.CountTotalTransactions `json:"total"`
}

func GetListTotalCountTransaction(c *gin.Context) {
	claims, err := util.Authorization(c)
	if err != nil {
		fmt.Println(err.Error())
		util.Response(c, 401, "Unauthorized", nil)
		return
	}

	trans, err := dao.GetCountTotalTransactions(claims["id"].(string))
	if err != nil {
		util.Response(c, 400, err.Error(), nil)
		return
	}

	var data totalCountTransactions
	data.Income = trans["Income"]
	data.Expence = trans["Expence"]

	var total dao.CountTotalTransactions
	total.TypeName = "Total"
	total.TotalAmount = trans["Income"].TotalAmount - trans["Expence"].TotalAmount
	data.Total = total

	util.Response(c, 200, "succes", data)
}
