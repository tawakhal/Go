package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func main() {
	initial := App{}
	initial.Initializer(
		os.Getenv("root"),
		os.Getenv(""),
		os.Getenv("dbpenerbit"),
	)
	initial.Run(":8888")
	fmt.Println("Aplikasi berjalan di port : 8888")
}

func (initial *App) Initializer(username, password, dbname string) {
	var err error
	ad := username + ":" + password + "@/" + dbname
	fmt.Println("Url Koneksi = ", ad)
	initial.DB, err = sql.Open("mysql", "root:@/dbpenerbit")
	//initial.DB, err = sql.Open("mysql", username+ ":"+ password+ "@/"+dbname)
	if err != nil {
		log.Fatal(err)
	}
	initial.Router = mux.NewRouter()
	initial.InitialRoute()
}

func (initial *App) InitialRoute() {
	initial.Router.HandleFunc("/", initial.getPengarang).Methods("GET")
}

func (initial *App) getPengarang(w http.ResponseWriter, r *http.Request) {
	p := Pengarang{}
	if err := p.getPeng(initial.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "No Record Found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJson(w, http.StatusOK, p)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJson(w, code, map[string]string{"error": message})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	respond, _ := json.Marshal(payload)
	w.Header().Set("Content-Type ", "application/json")
	w.WriteHeader(code)
	w.Write(respond)
}

type Pengarang struct {
	Kd_pengarang string `json:"Kd_pengarang"`
	Nama         string `json:"Nama"`
	Alamat       string `json:"Alamat"`
	Kota         string `json:"Kota"`
}

func (initial *App) Run(strAddress string) {
	log.Fatal(http.ListenAndServe(strAddress, initial.Router))
}
