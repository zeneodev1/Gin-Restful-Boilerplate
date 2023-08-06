package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/zeneodev1/gin-restful-boilerplate/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type command struct {
	cmd    string
	help   string
	action func()
}

var commands []command = []command{
	{
		cmd:    "create",
		help:   "Creates the database if doesn't exit",
		action: createDB,
	},
	{
		cmd:    "drop",
		help:   "Drops the database if exits",
		action: dropDB,
	},
	{
		cmd:    "reset",
		help:   "Drops if exists and creates the database",
		action: resetDB,
	},
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Error: No Arg is provided")
		os.Exit(1)
	}

	config.LoadConfig()

	if os.Args[1] == "help" {
		help()
		return
	}

	cmd, err := fetchCommand()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		help()
		os.Exit(1)
	}
	cmd.action()
}

func fetchCommand() (*command, error) {
	for _, c := range commands {
		if c.cmd == os.Args[1] {
			return &c, nil
		}
	}

	return nil, errors.New("command not found")
}

func createDB() {
	db := SetupDB()
	dbname := config.GetDBConfig().Name

	var count int
	db.Raw("SELECT count(*) FROM pg_database WHERE datname = ?;", dbname).Scan(&count)
	if count > 0 {
		fmt.Println("Database already exists")
		return
	}

	tx := db.Exec("CREATE DATABASE " + dbname + ";")
	if tx.Error != nil {
		fmt.Println("Error: ", tx.Error)
		return
	}

	fmt.Printf("Database '%s' is created successfully\n", dbname)
}

func dropDB() {
	db := SetupDB()
	dbname := config.GetDBConfig().Name

	var count int
	db.Raw("SELECT count(*) FROM pg_database WHERE datname = ?;", dbname).Scan(&count)
	if count == 0 {
		fmt.Println("Database doesn't exit")
		return
	}

	tx := db.Exec("DROP DATABASE " + dbname + ";")
	if tx.Error != nil {
		fmt.Println("Error: ", tx.Error)
		return
	}

	fmt.Printf("Database '%s' is dropped successfully\n", dbname)
}

func resetDB() {
	dropDB()
	createDB()
}

func SetupDB() *gorm.DB {
	config := config.GetDBConfig()
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=disable", config.User, config.Password, config.Host, config.Port)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return database
}

func help() {
	fmt.Printf(`
This cli is a tool for managing database creations and deleteing

Usage: 
	
	go run ./cmd/cli/db <command>

The commands are:

`)

	for _, c := range commands {
		fmt.Print("	")
		fmt.Printf("%-10s %s \n", c.cmd, c.help)
	}
	fmt.Println()
}
