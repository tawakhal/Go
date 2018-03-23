package main

import (
	rt "DBSellingGorilla/item/routers"
	"log"
	"net/http"
)

func main() {
	//Membuat variabel router
	router := rt.InitRouter()

	log.Fatal(http.ListenAndServe(":8888", router))
}
