package model

import "github.com/gin-gonic/gin"

type DataIN struct {
	GinContext *gin.Context
	UserID     string
	UserName   string
}