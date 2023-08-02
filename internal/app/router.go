package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zeneodev1/gin-restful-boilerplate/config"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/app/controllers"
)

func RunServer() {
	router := gin.Default()

	// routes configured here

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Message": "Hello world",
		})
	})

	var userCtrl controllers.UserCtrl = *controllers.NewUserCtrl()
	users := router.Group("/users")
	{
		users.GET("/", userCtrl.Index)
		users.GET("/:id", userCtrl.Show)
		users.POST("/", userCtrl.Create)
		users.PUT("/:id", userCtrl.Update)
		users.DELETE("/:id", userCtrl.Delete)
	}

	err := router.Run(addr())
	if err != nil {
		panic(err)
	}
}

func addr() string {
	config := config.GetServerConfig()
	return fmt.Sprintf("%v:%v", config.Host, config.Port)
}
