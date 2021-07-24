package api

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goflow/models"
	"goflow/pkg/app"
	"goflow/pkg/e"
	"goflow/pkg/util"
	"net/http"
)

type user struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func CreateUser(c *gin.Context)  {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	var a user
	if err := c.ShouldBindBodyWith(&a, binding.JSON); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	}

	ok, err := valid.Valid(&a)

	if !ok {
		fmt.Println(111111, err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	user := &models.User{Username: a.Username, Password: a.Password}

	err = user.CreateUser()
	if err == nil {
		appG.Response(http.StatusOK, e.ERROR_EXIST_USERNAME, nil)
		return
	}

	err = user.CreateUser()
	if err != nil {
		fmt.Println(22222222, err)
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	token, err := util.GenerateToken(user.Username, user.Password)
	if err != nil {
		fmt.Println(33333333, err)
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}

func GetUser(c *gin.Context)  {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	var a user
	if err := c.ShouldBindBodyWith(&a, binding.JSON); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	ok, err := valid.Valid(&a)

	if !ok {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	fmt.Println(a)
	user := &models.User{Username: a.Username, Password: a.Password}

	err = user.CheckUser()
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	token, err := util.GenerateToken(user.Username, user.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
