package server

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/notes", getAllNotes)
	router.GET("/note/:id", getNote)
	router.POST("/note", postNote)
	router.DELETE("/note/:id", deleteNoteById)

	return router
}
