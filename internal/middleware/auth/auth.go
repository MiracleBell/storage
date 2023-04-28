package auth

import (
	. "../../models/users"
	"net/http"
)

func LoginService(user User) bool {
	//TODO: сделать связь с БД

	if "" == user.UserName {
		return false
	}

	return true
}

func LogoutService() bool {
	//TODO: сделать связь с БД
	return false
}

func IsAuth(cookie *http.Cookie) bool {
	//TODO: сделать связь с БД
	return false
}
