package routers

import (
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {

	// Membuat variabel baru router dgn isi *mux.Router
	router := mux.NewRouter().StrictSlash(false)

	// Mengisi nilai Officer Router
	router = setOfficerRouter(router)

	return router
}
