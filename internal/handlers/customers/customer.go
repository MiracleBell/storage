package customers

import (
	. "../../services/customers"
	"encoding/json"
	"net/http"
)

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	var id uint64
	json.NewDecoder(r.Body).Decode(&id)
	user, err := GetById(id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(user)
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	users, err := GetAllCustomers()
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(users)
}
