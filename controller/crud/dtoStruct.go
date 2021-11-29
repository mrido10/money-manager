package crud

import "github.com/gin-gonic/gin"

type DTO struct {
	GinContext *gin.Context
	UserID     string
}