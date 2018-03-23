package main

import (
	rt "DBSellingGorilla/selling/routers"
	"log"
	"net/http"
)

func main() {
	//Membuat variabel router
	router := rt.InitRouter()

	// Informasi Fatal Error jika ada error dan melakukan set http
	log.Fatal(http.ListenAndServe(":7000", router))
}
