package fixtures

import (
	"github.com/zeneodev1/gin-restful-boilerplate/internal/app/models"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/app/repositories"
	"gorm.io/gorm"
)

func UserFixture() *models.User {
	user := &models.User{
		FullName: "some name",
		Email:    "some@email.com",
	}
	repositories.DB.Create(&user)

	return user
}

func ClearUsers() {
	repositories.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.User{})
}
