package hitung

import "fmt"

type Mahasiswa struct {
	Nama string
}

// Menggunakan huruf besar maka bersifat public yang bisa diakses disemua project
func Tambah(a, b int) {
	var hasil int
	hasil = a + b
	fmt.Println("Hasil Penjumlahan\t= ", hasil)
}

// Menggunakan huruf besar maka bersifat public yang bisa diakses disemua project
func Kali(a, b int) {
	var hasil int
	hasil = a * b
	fmt.Println("Hasil Penjumlahan\t= ", hasil)
}

// Menggunakan huruf kecil maka bersifat private yang bisa diakses disemua project
func bagi(a, b int) {
	var hasil int
	hasil = a / b
	fmt.Println("Hasil Pembagian\t= ", hasil)
}
