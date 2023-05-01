package auth

import (
	. "../../models/users"
	. "../../utils/db"
	. "../../utils/generator"
	"log"
	"net/http"
)

func LoginService(user User) string {
	var db, err = OpenConnection()
	var uuid = UuidGenerator()

	query := "UPDATE users SET cookie=? WHERE login=?"
	_, err = db.Exec(query, uuid, user.UserName)
	if err != nil {
		log.Fatal(err)
	}

	CloseConnection(db)
	return uuid
}

func LogoutService(cookie *http.Cookie) bool {
	var db, err = OpenConnection()
	query := "DELETE cookie FROM users WHERE cookie=?"
	_, err = db.Exec(query, cookie.Value)
	if err != nil {
		log.Fatal(err)
	}
	return true
}

func IsAuth(cookie *http.Cookie) bool {
	var db, err = OpenConnection()
	query := "SELECT cookie FROM users WHERE cookie=?"
	_, err = db.Exec(query, cookie)
	if err != nil {
		log.Fatal("Error querying database:", err)
		return false
	}
	return true
}

func RequireTokenAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/login" {
			cookie, err := r.Cookie("token")
			if err != nil {
				// Если куки не найдены, возвращаем ошибку
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				http.Redirect(w, r, "/login", http.StatusUnauthorized)
				return
			}
			if !IsAuth(cookie) {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		}
	})
}
