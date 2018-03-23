package controller

import (
	"database/sql"
	"day12/sellingGorilla/repositories"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (initial *App) Run(strAddress string) {
	log.Fatal(http.ListenAndServe(":8888", initial.Router))
}
func (intial *App) Initializer(username, password, dbname string) {
	var err error
	intial.DB, err = sql.Open("mysql", "root:root@/DBSelling")
	if err != nil {
		log.Fatal(err)
	}
	intial.Router = mux.NewRouter()
	intial.InitialRoute()
}

func (initial *App) InitialRoute() {
	initial.Router.HandleFunc("/item", initial.GetItem).Methods("Get")
}

func (initial *App) GetItem(w http.ResponseWriter, r *http.Request) {
	i := repositories.Item{}
	if err := i.GetItem(initial.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "No Record Found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJson(w, http.StatusOK, i)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJson(w, code, map[string]string{"error": message})
}
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	respond, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(code)
	w.Write(respond)

}
