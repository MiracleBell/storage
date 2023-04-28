package users

import (
	. "../../models/users"
	checker "../../utils"
	"encoding/json"
	"net/http"
)

var isAuth bool

func Login(w http.ResponseWriter, r *http.Request) {
	if !checker.IsPostMethod(r) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if !isAuth {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
	return
}

func Logout(w http.ResponseWriter, r *http.Request) {
	if !checker.IsDeleteMethod(r) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	_, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	return
}
