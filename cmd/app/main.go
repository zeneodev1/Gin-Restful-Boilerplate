package main

import (
	"fmt"

	"github.com/zeneodev1/gin-restful-boilerplate/config"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/app"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/app/repositories"
)

func main() {
	config.LoadConfig()
	repositories.ConnectDB()

	router := app.SetupRouter()
	err := router.Run(server_addr())
	if err != nil {
		panic(err)
	}
}

func server_addr() string {
	config := config.GetServerConfig()
	return fmt.Sprintf("%v:%v", config.Host, config.Port)
}
