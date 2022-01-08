package util

import (
	"money-manager/model/responseModel"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

var (
	resp responseModel.Response
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
	resp.Code = statusCode
	if statusCode == 200 {
		resp.Status = true
	} else {
		resp.Status = false
	}

	c.JSON(statusCode, resp)
}

func ToShowResp(data interface{}, err error, c *gin.Context) {
	if err != nil {
		log.Error(err.Error())
		Response(c, 400, err.Error(), nil)
		return
	}

	Response(c, 200, "succes", data)
}

func ResponseNotFound(c *gin.Context) {
	Response(c, http.StatusNotFound, "Not Found", nil)
}