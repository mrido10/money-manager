package controller

import (
	"money-manager/util"

	"github.com/gin-gonic/gin"
)

func Verify(c *gin.Context) {
	_, err := util.Authorization(c)
	if err != nil {
		util.Response(c, 401, "Unauthorized", nil)
		return
	}
	util.Response(c, 200, "Ok", nil)
}
