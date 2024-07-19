package postgres

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

type Note struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Text  string `json:"text"`
}

var db *sql.DB

// Return new Postgresql db instance
func InitDB() error {
    var err error
    db, err = sql.Open("postgres", "host=db user=postgres password=123 dbname=test sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    schema := `
    CREATE TABLE IF NOT EXISTS notes (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        text TEXT NOT NULL
    );`
    _, err = db.Exec(schema)
    if err != nil {
        return err
    }

    fmt.Println("Connected to the database")
    return nil // Return nil if no errors
}

func scanNote(rows *sql.Rows) (Note, error) {
    var note Note
    err := rows.Scan(&note.ID, &note.Title, &note.Text)
    return note, err
}
