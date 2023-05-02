package customers

import (
	. "../../models/customer"
	. "../../services/customers"
	"encoding/json"
	"net/http"
)

func PostOrganization(w http.ResponseWriter, r *http.Request) {
	var organization Organization
	json.NewDecoder(r.Body).Decode(&organization)
	AddOrganization(&organization)
	w.WriteHeader(http.StatusCreated)
}
