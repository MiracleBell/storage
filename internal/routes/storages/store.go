package storages

import . "../../handlers/storages"
import "net/http"

func StoreRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		PostGoods(w, r)
	case "GET":
		GetGoods(w, r)
	default:
		http.Error(w, "Invalid request method", 405)
	}
}

func StoreByIdRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetGoodsById(w, r)
	case "PUT":
		PutGoods(w, r)
	default:
		http.Error(w, "Invalid request method", 405)
	}
}
