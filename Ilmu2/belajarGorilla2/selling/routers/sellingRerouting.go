package routers

import (
	con "DBSellingGorilla/selling/controllers"

	"github.com/gorilla/mux"
)

func setSellingRouter(router *mux.Router) *mux.Router {
	//Pembuatan EndPoint Selling
	router.HandleFunc("/selling", con.GetSelling).Methods("GET")

	//Mengembalikan nilai router
	return router
}
