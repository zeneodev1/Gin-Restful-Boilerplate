package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	env    string
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Port string
	Host string
}

type DBConfig struct {
	Name     string
	Port     string
	Host     string
	User     string
	Password string
}

var config Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var env string
	if env = os.Getenv("GOENV"); env == "" {
		env = "dev"
	}
	config.env = env

	parseConfig()
}

func GetDBConfig() DBConfig {
	return config.DB
}

func GetServerConfig() ServerConfig {
	return config.Server
}

func parseConfig() {
	config.DB = DBConfig{
		Name:     os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}
	config.Server = ServerConfig{
		Port: serverPort,
		Host: os.Getenv("SERVER_HOST"),
	}
}
