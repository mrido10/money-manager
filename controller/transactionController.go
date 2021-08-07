package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAllTrancsactionNote(c *gin.Context) {
	c.String(http.StatusOK, "testing oke")
}


