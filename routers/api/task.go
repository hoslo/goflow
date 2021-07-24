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

func GetTasks(c *gin.Context)  {
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

	tasks, err := models.QueryTasks(uint(id))
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"tasks" : tasks,
	})
}

func CreateTask(c *gin.Context)  {
	appG := app.Gin{C: c}

	token := c.Request.Header.Get("token")
	id, err := strconv.ParseUint(c.Query("id"), 10,64)
	if err != nil {
		fmt.Println(0000, err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}


	_, err = util.ParseUser(token)
	if err != nil {
		fmt.Println(1111, err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	task := &models.Task{}

	if err := c.ShouldBindBodyWith(&task, binding.JSON); err != nil {
		fmt.Println(2222, err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	_, err = models.AddTask(uint(id), task)
	if err != nil {
		fmt.Println(3333, err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"dag" : task,
	})
}

func DeleteTask(c *gin.Context)  {
	appG := app.Gin{C: c}

	task := models.Task{}

	if err := c.ShouldBindBodyWith(&task, binding.JSON); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	}

	if task.ID != 0 {
		_, err := models.DeleteTask(&task)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		}
	}else {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"dag" : task,
	})
}
