package app

import (
	"github.com/gin-gonic/gin"
	"github.com/zeneodev1/gin-restful-boilerplate/config"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/app/controllers"
)

func SetupRouter() *gin.Engine {
	if config.GetEnv() != config.EnvDev {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Message": "Hello world",
		})
	})

	var userCtrl controllers.UserCtrl = *controllers.NewUserCtrl()
	users := router.Group("/users")
	{
		users.GET("", userCtrl.Index)
		users.GET("/:id", userCtrl.Show)
		users.POST("", userCtrl.Create)
		users.PUT("/:id", userCtrl.Update)
		users.DELETE("/:id", userCtrl.Delete)
	}

	return router
}
