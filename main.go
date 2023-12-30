package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type note struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:text`
}

var notes = []note{
	{ID: "1", Title: "Blue Train", Text: "John Coltrane"},
	{ID: "2", Title: "Jeru", Text: "Gerry Mulligan"},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Text: "Sarah Vaughan"},
}

func getNotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, notes)
}

func postNote(c *gin.Context) {
	var newNote note

	if err := c.BindJSON(&newNote); err != nil {
		return
	}

	notes = append(notes, newNote)
	c.IndentedJSON(http.StatusOK, newNote)
}

func getNote(c *gin.Context) {
	id := c.Param(("id"))

	for _, note := range notes {
		if note.ID == id {
			c.IndentedJSON(http.StatusOK, note)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/notes", getNotes)
	router.GET("/notes/:id", getNote)
	router.POST("/notes", postNote)

	router.Run("localhost:8080")
}
