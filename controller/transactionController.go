package controller

import (
	"fmt"
	"money-manager/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTrancsactionNote(c *gin.Context) {
	claims, err := util.Authorization(c)
	if err != nil {
		fmt.Println(err.Error())
		util.Response(c, 401, "Unauthorized", nil)
		return
	}

	fmt.Println(".. wellcome .. ", claims["name"])

	c.String(http.StatusOK, "testing oke")
}
