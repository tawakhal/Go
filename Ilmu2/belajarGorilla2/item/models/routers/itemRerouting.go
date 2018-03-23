package routers

import (
	con "DBSellingGorilla/item/controllers"

	"github.com/gorilla/mux"
)

func setItemRouter(router *mux.Router) *mux.Router {
	//Pembuatan EndPoint Item

	router.HandleFunc("/item", con.GetItem).Methods("GET")

	return router
}
