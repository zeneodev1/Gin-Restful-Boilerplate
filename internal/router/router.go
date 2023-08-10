package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zeneodev1/gin-restful-boilerplate/config"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	if config.GetEnv() != config.EnvDev {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Message": "Hello world",
		})
	})

	setupUserRouter(router, db)

	return router
}
