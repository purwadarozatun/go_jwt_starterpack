package main

import (
	"company/myproject/databases"
	"company/myproject/deps"
	"fmt"
	"log"
	"os/user"
)

func init() {
	config, err := deps.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	databases.ConnectDB(&config)
}

func main() {
	databases.DB.AutoMigrate(&user.User{})
	fmt.Println("? Migration complete")
}
