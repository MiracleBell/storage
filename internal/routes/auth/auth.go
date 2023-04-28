package auth

import (
	. "../../handlers/auth"
	. "../../utils"
	"net/http"
)

func LoginRouter(w http.ResponseWriter, r *http.Request) {
	if IsPostMethod(r) {
		Login(w, r)
	}
	http.Error(w, "Invalid request method", 405)
}

func LogoutRouter(w http.ResponseWriter, r *http.Request) {
	if IsDeleteMethod(r) {
		Logout(w, r)
	}
	http.Error(w, "Invalid request method", 405)
}
