package handlers

import (
	u "../middleware/users"
	checker "../utils"
	"net/http"
)

func StartPage(w http.ResponseWriter, r *http.Request) {
	if !checker.IsGetMethod(r) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	cookie, err := r.Cookie("token")
	if err != nil || !u.IsAuth(cookie) {
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	return
}
