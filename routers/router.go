package routers

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"goflow/routers/api"
	"time"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	cors.DefaultConfig()
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "token"},
		AllowCredentials: false,
		AllowAllOrigins: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/create_user", api.CreateUser)
	r.POST("/get_user", api.GetUser)

	r.POST("/create_dag", api.CreateDAG)
	r.POST("/delete_dag", api.DeleteDAG)
	r.GET("/get_dags", api.GetDAGs)
	r.GET("/get_dag", api.GetDAG)

	r.POST("/create_httpoperator", api.CreateHttpOperator)
	r.POST("/delete_httpoperator", api.DeleteHttpOperator)
	r.GET("/get_httpoperators", api.GetHttpOperators)
	r.GET("/get_httpoperator", api.GetHttpOperator)

	r.POST("/create_task", api.CreateTask)
	r.POST("/delete_task", api.DeleteTask)
	r.GET("/get_tasks", api.GetTasks)

	r.GET("/get_task_instances", api.GetTaskInstances)

	r.POST("/test1", func(c *gin.Context) {
		fmt.Println("test1")
		c.JSON(200, gin.H{
			"ret": 200,
			"err": "",
			"data": "test1",
		})
	})

	r.POST("/test2", func(c *gin.Context) {
		fmt.Println("test2")
		c.JSON(200, gin.H{
			"ret": 200,
			"err": "",
			"data": "test2",
		})
	})

	r.POST("/test3", func(c *gin.Context) {
		fmt.Println("test3")
		c.JSON(200, gin.H{
			"ret": 200,
			"err": "",
			"data": "test3",
		})
	})

	return r
}