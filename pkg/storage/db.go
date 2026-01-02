package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./automation.db")
	if err != nil {
		return nil, err
	}

	statement := `
	CREATE TABLE IF NOT EXISTS profiles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		url TEXT UNIQUE,
		status TEXT,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(statement)
	return db, err
}

func IsDuplicate(db *sql.DB, url string) bool {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM profiles WHERE url=?)", url).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func SaveProfile(db *sql.DB, url string, status string) {
	db.Exec("INSERT OR IGNORE INTO profiles (url, status) VALUES (?, ?)", url, status)
}