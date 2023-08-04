package repositories

import (
	"fmt"

	"github.com/zeneodev1/gin-restful-boilerplate/config"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(postgres.Open(datasource()), &gorm.Config{
		Logger: setupLogger(),
	})
	if err != nil {
		panic(err)
	}

	DB = db
	DB.AutoMigrate(&models.User{})
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

func setupLogger() logger.Interface {
	var logLevel logger.LogLevel

	switch config.GetEnv() {
	case config.EnvDev:
		logLevel = logger.Error
	case config.EnvTest:
		logLevel = logger.Silent
	case config.EnvProd:
		logLevel = logger.Silent
	}

	return logger.Default.LogMode(logLevel)
}
