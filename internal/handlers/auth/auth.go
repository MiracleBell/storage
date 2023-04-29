package auth

import (
	. "../../middleware/auth"
	. "../../models/users"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	var uuid = LoginService(user)
	cookie := http.Cookie{
		Name:     "token",
		Value:    uuid,
		HttpOnly: true}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/", http.StatusFound)
	return
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	LogoutService(cookie)
}
