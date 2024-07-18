package main

import (
	server "notes-api/internal"
	"notes-api/pkg/postgres"
)

func main() {
	postgres.InitDB()

	router := server.SetupRouter()
	router.Run("localhost:8080")
}
