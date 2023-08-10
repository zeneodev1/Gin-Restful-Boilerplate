package main

import (
	"fmt"

	"github.com/zeneodev1/gin-restful-boilerplate/config"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/repositories"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/router"
)

func main() {
	config.LoadConfig()

	db, err := repositories.ConnectDB()
	if err != nil {
		panic(err)
	}

	router := router.SetupRouter(db)
	err = router.Run(server_addr())
	if err != nil {
		panic(err)
	}
}

func server_addr() string {
	config := config.GetServerConfig()
	return fmt.Sprintf("%v:%v", config.Host, config.Port)
}
