package customer

import . "../../handlers/customer"
import "net/http"

func CustomerRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		PostCustomer(w, r)
	case "GET":
		GetCustomers(w, r)
	default:
		http.Error(w, "Invalid request method", 405)
	}
}

func CustomerWithId(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		GetCustomerById(w, r)
	}
	http.Error(w, "Invalid request method", 405)
}
