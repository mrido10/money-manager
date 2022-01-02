package controller

// import (
// 	"fmt"
// 	"money-manager/dao"
// 	"money-manager/util"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// func RegisterTo(c *gin.Context) {
// 	claims, err := util.Authorization(c)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		util.Response(c, 401, "Unauthorized", nil)
// 		return
// 	}

// 	fmt.Println(".. wellcome .. ", claims["name"])

// 	action := c.Query("action")
// 	name := c.Query("name")

// 	switch action {
// 	case "toAccountGroup":
// 		toAccountGroup(name, c)
// 	case "toAccount":
// 		toAccount(c)
// 	case "toTransactionType":
// 		toTransactionType(name, c)
// 	case "toTransactionCategory":
// 		toTransactionCategory(name, c)
// 	case "userAccount":

// 	default:
// 		fmt.Printf("Cant use this action: %s", action)
// 	}
// }

// func toAccountGroup(name string, c *gin.Context) {
// 	err := dao.RegisterToAccountGroup(name)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		util.GinJsonResp(c, 400, err.Error(), false)
// 		return
// 	}

// 	util.GinJsonResp(c, 200, "succes", true)
// }

// func toAccount(c *gin.Context) {
// 	groupID, _ := strconv.Atoi(c.Query("groupID"))
// 	accName := c.Query("accountName")
// 	amount, _ := strconv.Atoi(c.Query("amount"))
// 	desc := c.Query("desc")

// 	if accName == "" {
// 		fmt.Print("accounName is empty")
// 		util.GinJsonResp(c, 400, "accounName is empty", false)
// 		return
// 	}

// 	err := dao.RegisterToAccount(groupID, accName, amount, desc)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		util.GinJsonResp(c, 400, err.Error(), false)
// 		return
// 	}

// 	util.GinJsonResp(c, 200, "succes", true)
// }

// func toTransactionType(name string, c *gin.Context) {
// 	err := dao.RegisterToTransactionType(name)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		util.GinJsonResp(c, 400, err.Error(), false)
// 		return
// 	}

// 	util.GinJsonResp(c, 200, "succes", true)
// }

// func toTransactionCategory(name string, c *gin.Context) {
// 	err := dao.RegisterToTransactionCategory(name)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		util.GinJsonResp(c, 400, err.Error(), false)
// 		return
// 	}

// 	util.GinJsonResp(c, 200, "succes", true)
// }
