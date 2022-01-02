package util

import (
	"fmt"
	"money-manager/config"
	"money-manager/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type DataIn struct {
	model.DataIN
}

func Authorization(c *gin.Context) (jwt.MapClaims, error) {
	tokenString := c.Request.Header.Get("Authorization")
	conf, err := config.GetConfig()
	if err != nil {
		return nil, nil
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("SIGNING METHOD INVALID")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("SIGNING METHOD INVALID")
		}

		return []byte(conf.Jwt.Key), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func (e DataIn) JWTValidations(c *gin.Context, serve func(model.DataIN) (interface{}, int, string, error)) {
	claims, err := Authorization(c)
	if err != nil {
		Response(c, 401, "Unauthorized", nil)
		return
	}

	e.GinContext = c
	e.UserID = claims["id"].(string)
	e.UserName = claims["name"].(string)

	respData, respCode, respMsg, err := serve(e.DataIN)

	if err != nil {
		Response(c, 400, err.Error(), nil)
		return
	}

	if (respCode == 0) {
		respCode = 400
	}


	if respCode == 200 && respMsg == "" {
		respMsg = "success"
	}

	Response(c, respCode, respMsg, respData)
}
