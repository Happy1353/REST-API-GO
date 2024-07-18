package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"notes-api/pkg/postgres"
)

type Note struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

func ScanNote(rows *sql.Rows) (Note, error) {
	var note Note
	err := rows.Scan(&note.ID, &note.Title, &note.Text)
	return note, err
}

func GetNote(c *gin.Context) {
	id := c.Param("id")

	// convert id to number
	noteID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// get note with id
	note, err := postgres.GetNoteByID(noteID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Note not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, note)
}

func PostNote(c *gin.Context) {
	var newNote postgres.Note

	if err := c.BindJSON(&newNote); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	newID, err := postgres.CreateNote(newNote)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"id": newID})
}

func GetAllNotes(c *gin.Context) {
	notes, err := postgres.GetAllNotes()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve notes"})
		return
	}

	// Если notes == nil, можно вернуть пустой массив вместо nil
	if notes == nil {
		notes = []postgres.Note{}
	}

	c.IndentedJSON(http.StatusOK, notes)
}

func DeleteNoteById(c *gin.Context) {
	id := c.Param("id")
	// convert id to number
	noteID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = postgres.DeleteNoteById(noteID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
		return
	}
}
