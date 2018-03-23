package main

import "Penting/hitung"

// Menggunakan huruf besar maka bersifat public yang bisa diakses disemua project
func main() {
	var a = hitung.Mahasiswa{}
	a.Nama = ""

	hitung.Tambah(2, 2)
	hitung.Kali(2, 3)
}
