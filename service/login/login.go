package login

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"money-manager/config"
	"money-manager/dao"
	"money-manager/model"
	"money-manager/util"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type user struct {
	Email 		string `json:"email"`
	Password 	string `json:"password"`
	*gin.Context
}

type responseBody struct {
	Message string `json:"message"`
	Status bool `json:"status"`
	Data struct {
		Authorization string `json:"authorization"`
	} `json:"data"`
}

func Login(c model.DataIN) (data interface{}, responseCode int, msg string, err error) {
	var us user
	us.Context = c.GinContext
	user, err := us.setUser()
	if err != nil {
		log.Error(err.Error())
		return
	}

	jsonBody, _ := json.Marshal(user) 

	conf, _ := config.GetConfig()
	resp, err := http.Post(fmt.Sprintf("%s/login", conf.AuthServer.Ip), "application/json; charset=utf-8", bytes.NewBuffer(jsonBody))

	if err != nil {
		log.Error(err.Error())
		return
	}

	defer resp.Body.Close()

	responseCode = resp.StatusCode
	if responseCode != 200 {
		if responseCode == 401 {
			err = errors.New("wrong Email or Password")
		} else {
			err = errors.New("unkonown Error")
		}
		log.Error(err.Error())
		return
	}

	respBody := convertStringBodyToJSON(resp)

	data = respBody.Data
	us.Context.Request.Header.Set("Authorization", respBody.Data.Authorization)
	err = us.insertUserAccountIfNotExist()
	return
}

func (us user) setUser() (data user, err error) {
	if err = us.ShouldBindJSON(&data); err != nil {
		return
	}
	return
}

func convertStringBodyToJSON(resp *http.Response) (respBody responseBody) {
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &respBody)
	return
}

func (us user) insertUserAccountIfNotExist() (err error) {
	claims, err := util.Authorization(us.Context)
	if err != nil {
		return
	}

	count, err := dao.CheckUserAccount(claims["id"].(string))
	if err != nil {
		return
	}
	if count < 1 {
		err = dao.InsertUserAccount(claims["id"].(string), claims["name"].(string))
		if err != nil {
			return
		}
	}
	return nil
}