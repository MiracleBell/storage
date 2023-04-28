package handlers

import (
	. "../middleware/users"
	. "../utils"
	"net/http"
)

func PostCustomer(w http.ResponseWriter, r *http.Request) {

}

func GetCustomerById(w http.ResponseWriter, r *http.Request) {

}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	if !IsGetMethod(r) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cookie, err := r.Cookie("token")
	if err != nil || !IsAuth(cookie) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

}
