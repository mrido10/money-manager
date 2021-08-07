package util

import (
	"money-manager/model"

	"github.com/gin-gonic/gin"
)

var (
	resp model.Response
)

func GinJsonResp(c *gin.Context, code int, msg interface{}, status bool) {
	c.JSON(code, gin.H{
		"message": msg,
		"status":  status,
	})
}

func Response(c *gin.Context, statusCode int, msg string, data interface{}) {
	resp.Message = msg
	resp.Data = data
	if statusCode == 200 {
		resp.Status = true
	} else {
		resp.Status = false
	}

	c.JSON(statusCode, resp)
}
