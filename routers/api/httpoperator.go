package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goflow/models"
	"goflow/pkg/app"
	"goflow/pkg/e"
	"goflow/pkg/util"
	"net/http"
	"strconv"
)

func CreateHttpOperator(c *gin.Context)  {
	appG := app.Gin{C: c}

	token := c.Request.Header.Get("token")

	user, err := util.ParseUser(token)
	if err != nil {
		fmt.Println(1111, err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	HttpOperator := &models.HttpOperator{}

	if err := c.ShouldBindBodyWith(&HttpOperator, binding.JSON); err != nil {
		fmt.Println(2222, err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	HttpOperator.UserId = user.ID
	HttpOperator.Method = "POST"
	HttpOperator, err = models.AddHttpOperator(HttpOperator)
	if err != nil {
		fmt.Println(3333, err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"httpoperator" : HttpOperator,
	})
}

func GetHttpOperators(c *gin.Context)  {
	appG := app.Gin{C: c}

	token := c.Request.Header.Get("token")

	user, err := util.ParseUser(token)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	HttpOperators, err := models.QueryHttpOperators(user.ID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"httpoperators" : HttpOperators,
	})
}

func GetHttpOperator(c *gin.Context)  {
	appG := app.Gin{C: c}

	token := c.Request.Header.Get("token")
	id, err := strconv.ParseUint(c.Query("id"), 10,64)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	_, err = util.ParseUser(token)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	HttpOperator, err := models.QueryHttpOperator(uint(id))
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"httpoperator" : HttpOperator,
	})
}

func DeleteHttpOperator(c *gin.Context)  {
	appG := app.Gin{C: c}

	HttpOperator := models.HttpOperator{}

	if err := c.ShouldBindBodyWith(&HttpOperator, binding.JSON); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	}

	if HttpOperator.ID != 0 {
		_, err := models.DeleteHttpOperator(&HttpOperator)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		}
	}else {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"httpoperator" : HttpOperator,
	})
}
