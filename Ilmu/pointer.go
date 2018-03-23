package main

import "fmt"

func main() {
	pointer1()
}

/*
	Pointer adalah sebuah alamat dari suatu nilai
	contoh alamat dari int
	rumus :
		var <namaVariabel> *<tipeData> = &<NamaVariabel yang dituju>
		contoh:
			var angka int = 4 // variabel dengan nilai 4
			var alamat *int = &angka // Pointer dengan nilai dari "angka"

	Kelemahan pointer :
	1. Menyebabkan beratnya suatu proses jika menggunakan banyak pointer

	lebih baik pointer jika banyak dimasukan/dibungkus didalam array/map/slice
	daripada mendeklarasikan satu persatu
*/

func pointer1() {
	var (
		angka  int  = 4
		alamat *int = &angka // ini sebuah pointer
	)

	fmt.Println("Nilai dari angka\t= ", angka)         //Mencetak nilai dari angka
	fmt.Println("Alamat dari angka\t= ", &angka)       //Mencetak alamat dari angka
	fmt.Println("Nilai dari alamat\t= ", alamat)       //Mencetak nilai alamat
	fmt.Println("Nilai asli dari alamat\t= ", *alamat) //Mencetak nilai asli dari alamat == angka

	//Merubah angka yang tadinya 4 menjadi
	angka = 10
	fmt.Println("\n\n\nDATA SETELAH DIRUBAH MENJADI 10") //Mencetak nilai dari angka
	fmt.Println("Nilai dari angka\t= ", angka)           //Mencetak nilai dari angka
	fmt.Println("Alamat dari angka\t= ", &angka)         //Mencetak alamat dari angka
	fmt.Println("Nilai dari alamat\t= ", alamat)         //Mencetak nilai alamat
	fmt.Println("Nilai asli dari alamat\t= ", *alamat)   //Mencetak nilai asli dari alamat == angka

	/*
		Walaupun nilai dari sih angka berubah
		Nilai pointer akan tetap sama dengan yang awal
		(jika angka ada / digunakan / tidak dihapus / dibuat lagi variabel yang baru)
	*/
}
