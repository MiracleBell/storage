package orders

import (
	. "../../handlers/orders"
	"net/http"
)

func OrderRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetOrders(w, r)
	case "POST":
		PostOrder(w, r)
	default:
		http.Error(w, "Invalid request method", 405)
	}
	http.HandleFunc("/orders", PostOrder)
	http.HandleFunc("/orders/{id}", GetOrderById)
}

func OrderWithIdRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetOrderById(w, r)
	default:
		http.Error(w, "Invalid request method", 405)
	}
}
