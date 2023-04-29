package db

import (
	"database/sql"
	"log"
)

const m_DRIVER_NAME = "mysql"

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open(m_DRIVER_NAME, "user:password@tcp(localhost:3306)/mydatabase")
	if err != nil {
		log.Fatal("Error opening database:", err)
		return nil, err
	}
	return db, nil
}

func CloseConnection(db *sql.DB) {
	defer db.Close()
}
