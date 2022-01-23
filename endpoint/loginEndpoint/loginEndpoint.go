package loginEndpoint

import (
	"money-manager/service/login"
	"money-manager/util"
	"github.com/gin-gonic/gin"
)

type loginStruct struct {
	util.DataIn
}

var LoginEndpoint = loginStruct{}

func (e loginStruct) Login(c *gin.Context) {
	e.ServeWithoutJWTValidations(c, login.Login)
}