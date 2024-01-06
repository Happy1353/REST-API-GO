package main

import (
	"notes-api/internal/server"
	"notes-api/pkg/postgres"
)

func main() {
	postgres.InitDB1()

	router := server.SetupRouter()
	router.Run("localhost:8080")
}
