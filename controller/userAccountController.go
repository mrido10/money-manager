package controller

// import (
// 	"fmt"
// 	"money-manager/dao"
// 	"money-manager/model/domain"
// 	"money-manager/util"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// type (
// 	UserAccount domain.UserAccount
// )

// func SaveUserAccount(c *gin.Context) {
// 	claims, err := util.Authorization(c)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		util.Response(c, 401, "Unauthorized", nil)
// 	}

// 	id, name := claims["id"].(string), claims["name"].(string)
// 	now := time.Now()
// 	var acc = UserAccount{id, name, now}

// 	// if err := c.ShouldBindJSON(&acc); err != nil {
// 	// 	fmt.Println(err)
// 	// 	util.Response(c, 400, err.Error(), nil)
// 	// 	return
// 	// }

// 	acc.SaveUser(c)
// }

// func (acc UserAccount) SaveUser(c *gin.Context) {
// 	countAcc, err := dao.CheckUserAccount(acc.UserID)

// 	if err != nil {
// 		util.Response(c, 400, err.Error(), nil)
// 		return
// 	}

// 	if countAcc == 0 {
// 		err := dao.InsertUserAccount(acc.UserID, acc.UserName)
// 		if err != nil {
// 			util.Response(c, 400, err.Error(), nil)
// 			return
// 		}

// 		util.Response(c, 200, "Succes insert new userAccount", nil)
// 		return
// 	}

// 	util.Response(c, 200, "userAccount have been exist", nil)
// }
