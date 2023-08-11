package helpers

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/zeneodev1/gin-restful-boilerplate/config"
	"github.com/zeneodev1/gin-restful-boilerplate/internal/repositories"
	"gorm.io/gorm"
)

func SetupEnv() {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Join(filepath.Dir(filename), "../..")
	os.Chdir(dir)
	os.Setenv("GOENV", "test")
	config.LoadConfig()
}

const savePoint string = "BEGINNING"

func SetupTx() (*gorm.DB, error) {
	db, err := repositories.ConnectDB()
	if err != nil {
		return nil, err
	}

	tx := db.Begin()
	tx = tx.SavePoint(savePoint)
	return tx, nil
}

func StartOverTx(tx *gorm.DB) *gorm.DB {
	tx = tx.RollbackTo(savePoint)

	return tx
}

func Rollback(tx *gorm.DB) {
	tx.Rollback()
}
