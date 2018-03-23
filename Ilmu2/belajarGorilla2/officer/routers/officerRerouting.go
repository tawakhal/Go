package routers

import (
	con "DBSellingGorilla/officer/controllers"

	"github.com/gorilla/mux"
)

func setOfficerRouter(router *mux.Router) *mux.Router {
	//Pembuatan EndPoint Officer
	router.HandleFunc("/officer", con.GetOfficer).Methods("GET")

	//Mengembalikan nilai router
	return router
}
