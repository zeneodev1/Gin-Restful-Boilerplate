package repositories

import (
	"fmt"

	"github.com/zeneodev1/gin-restful-boilerplate/config"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(datasource()), &gorm.Config{
		Logger: setupLogger(),
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})

	return db, nil
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
