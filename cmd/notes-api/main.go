package main

import (
	"fmt"
	server "notes-api/internal"
	"notes-api/pkg/postgres"
)

func main() {
	err := postgres.InitDB()
	if err != nil {
		fmt.Println("Failed to initialize database:", err)
		return
	}

	router := server.SetupRouter()
	router.Run(":8080")
}
