package main

import (
	"flourish-coding-challenge/cmd"
	"flourish-coding-challenge/config"
	"flourish-coding-challenge/internal/db"
	"fmt"
	"os"
)

func init() {
	config.LoadEnv(".env")
	dsn := os.Getenv("DB_STRING")
	if err := db.ConnectDB(dsn); err != nil {
		fmt.Println("Failed to connect to the database:", err)
		os.Exit(1)
	}
	db.Migrate()

}

func main() {
	cmd.Execute()

}
