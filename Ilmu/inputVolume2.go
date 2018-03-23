package main

import (
	"fmt"
	"math"
)

var (
	mau        int
	pilih      int
	p, l, r, t float32
)

func main() {

	for mau == 1 {
		fmt.Print("Masukan volume apa yang anda ingin hitung\n1.Balok\n2.Tabung\n3.Kerucut\n\nPilihan Anda :  ")
		//Untuk mendapatkan nilai kedalam variabel pilih dengan tipe data Integer dari inputan keyboard
		fmt.Scanln(&pilih)
		switch pilih {
		case 1:
			cetakBalok()
			fallthrough
		case 2:
			cetakTabung()
		case 3:
			cetakKerucut()
		default:
			fmt.Println("Tolong masukan Pilihan yang benar")

		}

		fmt.Print("\n\n\nApakah anda ingin menghitung volume ?\n1.Ya\n2.Tidak\nMasukan Pilihan Anda :  ")
		fmt.Scanln(&mau)
		fmt.Println("")
	}

}

func cetakBalok() {
	fmt.Print("\nMasukan Panjang Balok	= ")
	fmt.Scanln(&p)
	fmt.Print("\nMasukan Lebar Balok	= ")
	fmt.Scanln(&l)
	fmt.Print("\nMasukan Tinggi Balok	= ")
	fmt.Scanln(&t)
	fmt.Print("\nVolume Balok  adalah	= ", p*l*t)
}

func cetakTabung() {
	fmt.Print("\nMasukan Jari - jari Tabung	= ")
	fmt.Scanln(&r)
	fmt.Print("\nMasukan tinggi Tabung		= ")
	fmt.Scanln(&t)
	fmt.Print("\nVolume Tabung adalah	= ", math.Pi*r*r*t)
}

func cetakKerucut() {
	fmt.Print("\nMasukan Jari - jari Kerucut	= ")
	fmt.Scanln(&r)
	fmt.Print("\nMasukan tinggi Kerucut		= ")
	fmt.Scanln(&t)
	fmt.Print("\nVolume Kerucut adalah	= ", (math.Pi*r*r*t)/3)
}
