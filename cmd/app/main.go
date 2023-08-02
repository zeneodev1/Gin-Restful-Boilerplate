package main

import (
	"github.com/zeneodev1/gin-restful-boilerplate/config"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/app"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/app/repositories"
)

func main() {
	config.LoadConfig()
	repositories.ConnectDB()
	app.RunServer()
}
