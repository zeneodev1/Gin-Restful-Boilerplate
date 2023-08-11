package fixtures

import (
	"github.com/zeneodev1/gin-restful-boilerplate/internal/models"
	"gorm.io/gorm"
)

func UserFixture(db *gorm.DB) *models.User {
	user := &models.User{
		FullName: "some name",
		Email:    "some@email.com",
	}
	db.Create(&user)

	return user
}
