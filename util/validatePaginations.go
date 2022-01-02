package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func IsExistParameterLimitAndOffset(c *gin.Context) (bool, string, int, int) {
	limit := c.Query("limit")
	offset := c.Query("offset")

	if limit == "" || offset == "" {
		return false, "paremeter 'limit' or 'offset' can't empty", 0, 0
	}

	intLimit, err1 := strconv.Atoi(limit)
	intOffset, err2 := strconv.Atoi(limit)
	if err1 != nil || err2 != nil {
		return false, "parameter 'limit' or 'offset' must be integer", 0, 0
	}
	return true, "", intLimit, intOffset
}