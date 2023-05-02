package customers

import (
	. "../../models/customer"
	. "../../services/customers"
	"encoding/json"
	"net/http"
)

func PostIndividual(w http.ResponseWriter, r *http.Request) {
	var individual Individual
	json.NewDecoder(r.Body).Decode(&individual)
	AddIndividual(&individual)
	w.WriteHeader(http.StatusCreated)
}
