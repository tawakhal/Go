package main

import (
	fmt "fmt"
)

func main() {
	//go routine
	//mesti dipanggil function
	//didepan fungsi di tambah go

	selesai := make(chan bool) // Membuat Channel
	go salam(selesai)          // dengan menambahkan kata 'Go' itu seudah menjadi Ansyncronus
	/*
		Cara ini yaitu disebut unbufferchannel dan ini Ansyncronus yang dimana dia akan ditaro kedalam locasi yang berbeda"
		Walaupun ini akan menjadi ketergantungan tapi dengan alamat yang berbeda jadi tidak menumpuk disatu tempat
	*/
	<-selesai //Ini dipanggil sesuai mau yang mana dulu yang selesai

	fmt.Println("Xsis Mitra Utama")
}

func salam(selesai chan bool) {
	fmt.Println("Selamat Datang diBatch142")
	selesai <- true
}

//Ini adalah cara pertama untuk menjalankan go routime tapi ada kelemahan yaitu waktu menunggu untuk eksekusi
func main2() {
	go salam2()
	//t.Sleep(1 * t.Second) bisa menjalankan go routime tapi menggunakan waktu selama 1 detik
	fmt.Println("Xsis Mitra Utama")
}

func salam2() {
	fmt.Println("Selamat Datang diBatch142")
}
