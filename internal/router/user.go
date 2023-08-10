package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/controllers"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/repositories"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/services"
	"gorm.io/gorm"
)

func setupUserRouter(router *gin.Engine, db *gorm.DB) {
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userCtrl := controllers.NewUserCtrl(userService)

	users := router.Group("/users")
	{
		users.GET("", userCtrl.Index)
		users.GET("/:id", userCtrl.Show)
		users.POST("", userCtrl.Create)
		users.PUT("/:id", userCtrl.Update)
		users.DELETE("/:id", userCtrl.Delete)
	}
}
