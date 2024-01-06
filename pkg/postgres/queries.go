package postgres

import (
	"database/sql"
	"fmt"
	"log"
)

// Получить запись из базы данных по ID
func GetNoteByID(id int) (Note, error) {
	var note Note

	// Подготовка SQL-запроса
	query := "SELECT id, title, text FROM notes WHERE id = $1"
	rows, err := db.Query(query, id)
	if err != nil {
		log.Fatal(err)
		return note, err
	}
	defer rows.Close()

	// Получение результатов запроса
	if rows.Next() {
		note, err := scanNote(rows)
		if err != nil {
			log.Fatal(err)
			return note, err
		}
		return note, nil
	}

	// Если запись не найдена
	return note, sql.ErrNoRows
}

func CreateNote(newNote Note) (int, error) {
	// Подготовка SQL-запроса для вставки новой записи и получения её ID
	query := "INSERT INTO notes (title, text) VALUES ($1, $2) RETURNING id"
	row := db.QueryRow(query, newNote.Title, newNote.Text)

	var newID int
	err := row.Scan(&newID)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	fmt.Println("Created note with ID:", newID)
	return newID, nil
}

func GetAllNotes() ([]Note, error) {
	var notes []Note

	query := "SELECT id, title, text FROM notes"
	rows, err := db.Query(query)
	if err != nil {
		return notes, err
	}
	defer rows.Close()

	for rows.Next() {
		note, err := scanNote(rows)
		if err != nil {
			return notes, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func DeleteNoteById(id int) error {
	query := "DELETE FROM notes WHERE id = $1"
	_, err := db.Exec(query, id)
	fmt.Println("Delete note with ID:", id)
	return err
}
