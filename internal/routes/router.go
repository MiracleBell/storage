package routes

import (
	. "../middleware/auth"
	. "./auth"
	. "./customer"
	. "./employees"
	. "./goods"
	. "./orders"
	. "./storages"
	. "./tasks"
	"net/http"
)

func MainRouter() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle(
		"/",
		RequireTokenAuthentication(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})),
	)

	//http.HandleFunc("/", StartPage)

	http.HandleFunc("/login", LoginRouter)
	http.HandleFunc("/logout", LogoutRouter)

	http.HandleFunc("/customers", CustomerRouter)
	http.HandleFunc("/customers/{id}", CustomerWithId)

	http.HandleFunc("/goods", GoodsRouter)
	http.HandleFunc("/goods/{id}", GoodsWithIdRouter)

	http.HandleFunc("/store/{storeId}/orders", OrderRouter)
	http.HandleFunc("/store/{storeId}/orders/{id}", OrderWithIdRouter)

	http.HandleFunc("/employee", EmployeeRouter)
	http.HandleFunc("/employee/{id}", EmployeeWithIdRouter)

	http.HandleFunc("/store/{storeId}/goods", StoreRouter)
	http.HandleFunc("/store/{storeId}/goods/{goodsId}", StoreByIdRouter)

	http.HandleFunc("/store/{storeId}/tasks", TaskRouter)
	http.HandleFunc("/store/{storeId}/tasks/{taskId}", TaskWithIdRouter)

	// запускаем сервер
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
