package test_setup

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/zeneodev1/gin-restful-boilerplate/config"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/app/repositories"
)

func SetupEnv() {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Join(filepath.Dir(filename), "..")
	os.Chdir(dir)
	os.Setenv("GOENV", "test")
	config.LoadConfig()
}

func SetupRepo() {
	SetupEnv()
	repositories.ConnectDB()
}
