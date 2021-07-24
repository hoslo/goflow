package routers

import (
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

	r.POST("/create_task", api.CreateTask)
	r.POST("/delete_task", api.DeleteTask)
	r.GET("/get_tasks", api.GetTasks)

	return r
}