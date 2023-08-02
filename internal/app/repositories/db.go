package repositories

import (
	"fmt"

	"github.com/zeneodev1/gin-restful-boilerplate/config"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Repo *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(postgres.Open(datasource()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	Repo = db
	Repo.AutoMigrate(&models.User{})
}

func datasource() string {
	config := config.GetDBConfig()
	name := config.Name
	host := config.Host
	port := config.Port
	user := config.User
	password := config.Password
	return fmt.Sprintf("dbname=%s host=%s port=%s user=%s password=%s sslmode=disable", name, host, port, user, password)
}
