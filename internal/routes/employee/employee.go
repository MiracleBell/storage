package employee

import (
	. "../../handlers/emploee"
	"net/http"
)

func EmployeeRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetEmployee(w, r)
	case "POST":
		PostEmployee(w, r)
	default:
		http.Error(w, "Invalid request method", 405)
	}
}

func EmployeeWithIdRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetEmployeeById(w, r)
	case "PUT":
		PutEmployee(w, r)
	case "DELETE":
		DeleteEmployee(w, r)

	default:
		http.Error(w, "Invalid request method", 405)
	}
}
