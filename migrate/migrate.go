package main

import (
	"fmt"
	"log"
	"{project_package}/databases"
	"{project_package}/deps"
	"{project_package}/models"
)

func init() {
	config, err := deps.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	databases.ConnectDB(&config)
}

func main() {
	databases.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}
