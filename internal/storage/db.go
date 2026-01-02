package storage

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/glebarez/go-sqlite" 
)

var DB *sql.DB

// InitDB sets up the SQLite database and creates the history table
func InitDB() {
	var err error
	// The driver name registered by the glebarez package is "sqlite"
	DB, err = sql.Open("sqlite", "./bot_history.db") 
	if err != nil {
		log.Fatal("Connection Error:", err)
	}

	// Verify the connection is actually alive
	if err = DB.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}

	statement := `
	CREATE TABLE IF NOT EXISTS contacted_people (
		url TEXT PRIMARY KEY,
		contacted_at DATETIME
	);`
	_, err = DB.Exec(statement)
	if err != nil {
		log.Fatal("Table Creation Error:", err)
	}
}
// WasContacted checks if we have already messaged this person
func WasContacted(url string) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM contacted_people WHERE url=?)"
	err := DB.QueryRow(query, url).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

// MarkAsContacted saves the URL to the database
func MarkAsContacted(url string) {
	statement := "INSERT INTO contacted_people (url, contacted_at) VALUES (?, ?)"
	_, err := DB.Exec(statement, url, time.Now())
	if err != nil {
		log.Printf("Error saving to DB: %v", err)
	}
}
