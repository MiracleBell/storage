package customer

import (
	. "../../middleware/auth"
	"net/http"
)

func PostCustomer(w http.ResponseWriter, r *http.Request) {

}

func GetCustomerById(w http.ResponseWriter, r *http.Request) {

}

func GetCustomers(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("token")
	if err != nil || !IsAuth(cookie) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

}
