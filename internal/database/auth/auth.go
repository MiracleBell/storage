package auth

import (
	. "../"
	"log"
	"net/http"
)

var db, _ = OpenConnection()

func AddCookie(cookie string) error {
	insertQuery := "INSERT INTO users(cookie) VALUES (?)"
	_, err := db.Exec(insertQuery, cookie)
	if err != nil {
		log.Fatal("Can't insert cookie")
		return err
	}
	return nil
}

func RemoveCookie(cookie *http.Cookie) error {
	deleteQuery := "DELETE FROM users WHERE cookie=?"
	_, err := db.Exec(deleteQuery, cookie.Value)
	if err != nil {
		log.Fatal("Can't remove cookie")
		return err
	}
	return nil
}

func HasCookie(cookie *http.Cookie) bool {
	query := "SELECT * FROM users WHERE cookie=?"
	_, err := db.Exec(query, cookie.Value)
	if err != nil {
		return false
	}
	return true
}
