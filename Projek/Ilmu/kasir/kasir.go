package main

import (
	"fmt"
)

var (
	uang, jumlah, total int
	pilih               string
	kode_barang         = []string{"1", "2", "3"}
	nama_barang         = []string{"Indomie", "Kopi", "Teh"}
	harga_barang        = []int{3000, 1000, 1000}
)

func main() {
	fmt.Println("Selamat Datang di Toko Condoriano")
	fmt.Println("Apakah anda ingin memesan barang kami ?")
	fmt.Println("Daftar Barang")
	cetakDaftar()
	fmt.Println("=================================================")
	fmt.Print("Masukan Nomor atau nama barang	=  ")
	fmt.Scanln(&pilih)
	fmt.Println()

	if pilihan(pilih) == 1 {
		transaksi(1)
	} else if pilihan(pilih) == 2 {
		transaksi(2)
	} else if pilihan(pilih) == 3 {
		transaksi(3)
	} else {
		fmt.Println("Barang anda tidak tersedia")
	}
}

//	Metod mencetak Daftar barang yang dikeluarkan dari Array

func cetakDaftar() {
	for i := 0; i < 3; i++ {
		fmt.Println(kode_barang[i], " ", nama_barang[i], "	Rp. ", harga_barang[i])
	}
}

//	Metod untuk mengecek dan mengembalikan nilai berupa int yang nanti jadi penentu pemesanan
func pilihan(pilih string) int {
	switch {
	case pilih == "1" || pilih == "indomie" || pilih == "Indomie":
		return 1
	// case "indomie":
	// 	return 1
	case pilih == "2" || pilih == "kopi" || pilih == "Kopi":
		return 2
	// case "kopi":
	// 	return 1
	case pilih == "3" || pilih == "teh" || pilih == "Teh":
		return 3
	// case "teh":
	// 	return 1
	default:
		return 0
	}
}

//	Metod untuk proses transaksi pemesanan
func transaksi(angka int) {
	fmt.Print("Jumlah yang dibeli		=  ")
	fmt.Scanln(&jumlah)
	fmt.Println()
	fmt.Print("Masukan Jumlah Uang Anda	=  ")
	fmt.Scanln(&uang)
	fmt.Println()

	total = harga_barang[angka-1] * jumlah

	fmt.Print("anda Memesan ", jumlah, " buah, maka Total	=  ", total)
	fmt.Println()

	if uang < total {
		fmt.Println("Uang anda kurang sebesar 	=  Rp. ", uang-total)
	} else {
		fmt.Println("Kembalian anda adalah 		=  Rp. ", uang-total)
	}
}
