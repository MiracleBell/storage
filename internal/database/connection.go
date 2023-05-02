package database

import (
	. "../../config"
	"database/sql"
	"log"
)

const m_DRIVER_NAME = "mysql"

func OpenConnection() (*sql.DB, error) {
	var config = LoadConfigDBFromFile()
	db, err := sql.Open(m_DRIVER_NAME, ToString(config))
	if err != nil {
		log.Fatal("Error opening database:", err)
		return nil, err
	}
	defer db.Close()
	//TODO: будет ли так работать?
	return db, nil
}
