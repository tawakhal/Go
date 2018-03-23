package controllers

import (
	"database/sql"
	"fmt"
	"log"
)

type Context struct {
	DB *sql.DB
}

// Ini adalah Data Link Layar
func DBInitial(*sql.DB) *sql.DB {
	//Membuka Koneksi ke dbselling dan dimasukan kedalam db, dan err
	db, err := sql.Open("mysql", "toor:toor@/dbselling")

	//Mengecek error
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	//Mengembalikan nilai koneksi
	return db
}
