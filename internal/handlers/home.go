package handlers

import (
	. "../middleware/auth"
	"net/http"
)

func StartPage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil || !IsAuth(cookie) {
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
