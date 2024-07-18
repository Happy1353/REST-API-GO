package internal

import (
	handlers "notes-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	authorized := router.Group("/")
	authorized.Use(handlers.AuthMiddleware())
	{
		authorized.GET("/notes", handlers.GetAllNotes)
		authorized.GET("/note/:id", handlers.GetNote)
		authorized.POST("/note", handlers.PostNote)
		authorized.DELETE("/note/:id", handlers.DeleteNoteById)
	}

	return router
}
