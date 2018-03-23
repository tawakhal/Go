package controllers

import (
	"DBSellingGorilla/officer/data"
	"encoding/json"
	"net/http"
)

func GetOfficer(w http.ResponseWriter, r *http.Request) {
	// Membuat variabel baru context dari struct Context
	context := Context{}

	// Membuat variabel Koneksi yang isinya koneksi yang sudah dibuka *sql.DB
	koneksi := DBInitial(context.DB)

	// Menutup d close jika sudah tidak digunakan lg di akhir penggunaannya
	defer koneksi.Close()

	// Membuat variabel repo dan mengisi kedalam struct di officerRepository yaitu Koneksi (*sql.DB)
	repo := &data.OfficerRepository{koneksi}

	// Mengambil semua data dari officer
	officer := data.GetAll(repo)

	//Melakukan Set diheader di header
	w.Header().Set("Content-Type", "application")

	//Tampilkan status oke
	w.WriteHeader(http.StatusOK)

	// membuat variabel cetak dan mengubahnya menjadi Json
	cetak, _ := json.Marshal(officer)

	//Tampilkan datanya
	w.Write(cetak)
}
