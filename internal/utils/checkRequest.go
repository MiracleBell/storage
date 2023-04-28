package utils

import "net/http"

func IsPostMethod(r *http.Request) bool {
	return r.Method == "POST"
}

func IsGetMethod(r *http.Request) bool {
	return r.Method == "GET"
}

func IsPutMethod(r *http.Request) bool {
	return r.Method == "PUT"
}

func IsDeleteMethod(r *http.Request) bool {
	return r.Method == "DELETE"
}
