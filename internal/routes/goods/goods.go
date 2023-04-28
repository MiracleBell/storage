package goods

import (
	. "../../handlers/goods"
	"net/http"
)

func GoodsRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		PostGoods(w, r)
	case "GET":
		GetGoods(w, r)
	default:
		http.Error(w, "Invalid request method", 405)
	}
}

func GoodsWithIdRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetGoodsById(w, r)
	case "PUT":
		PutGoods(w, r)
	case "DELETE":
		DeleteGoods(w, r)
	default:
		http.Error(w, "Invalid request method", 405)
	}
}
