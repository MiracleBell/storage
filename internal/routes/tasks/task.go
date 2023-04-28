package tasks

import "net/http"
import . "../../handlers/tasks"

func TaskRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		PostTask(w, r)
	case "GET":
		GetTasks(w, r)
	default:
		http.Error(w, "Invalid request method", 405)
	}
}

func TaskWithIdRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		PutTasks(w, r)
	case "GET":
		GetTaskById(w, r)
	default:
		http.Error(w, "Invalid request method", 405)
	}
}
