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

func CreateDAG(c *gin.Context)  {
	appG := app.Gin{C: c}

	token := c.Request.Header.Get("token")

	user, err := util.ParseUser(token)
	if err != nil {
		fmt.Println(1111, err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	dag := &models.DAG{}

	if err := c.ShouldBindBodyWith(&dag, binding.JSON); err != nil {
		fmt.Println(2222, err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	dag, err = models.AddDAG(user.ID, dag)
	if err != nil {
		fmt.Println(3333, err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"dag" : dag,
	})
}

func GetDAGs(c *gin.Context)  {
	appG := app.Gin{C: c}

	token := c.Request.Header.Get("token")

	user, err := util.ParseUser(token)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	dags, err := models.QueryDAGs(user.ID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"dags" : dags,
	})
}

func GetDAG(c *gin.Context)  {
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

	dag, err := models.QueryDAG(uint(id))
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"dag" : dag,
	})
}

func DeleteDAG(c *gin.Context)  {
	appG := app.Gin{C: c}

	dag := models.DAG{}

	if err := c.ShouldBindBodyWith(&dag, binding.JSON); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	}

	if dag.ID != 0 {
		_, err := models.DeleteDAG(&dag)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		}
	}else {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"dag" : dag,
	})
}
