package main

import (
	"fmt"
	"strconv"
)

var (
	n int
)

func main() {
	// fmt.Println("SOAL NO 1 ")
	// soal1()
	// fmt.Println()
	//fmt.Println("SOAL NO 2 ")
	//soal2()
	// fmt.Println()
	// fmt.Println("SOAL NO 3 ")
	// soal3()
	// fmt.Println()
	fmt.Println("SOAL NO 4 ")
	soal4()
}

func soal1() {
	fmt.Print("Masukan jumlah n = ")
	fmt.Scanln(&n)

	//Membuat Baris
	pola := make([][]string, n)

	//Membuat kolom
	for i := range pola {
		pola[i] = make([]string, n)
	}

	var ganjil int = 1
	var genap int = (n * 2) - 2

	//Memasukan nilai kedalam array
	for x := 0; x < len(pola); x++ {
		for y := 0; y < len(pola); y++ {
			if x == y { //garis miring kanan kebawah
				pola[x][y] = strconv.Itoa(ganjil) + "\t"
				ganjil += 2
			} else if y+x == n-1 { //garis miring kanan kebawah
				pola[x][y] = strconv.Itoa(genap) + "\t"
			} else if y >= x && x+y <= n-1 {
				pola[x][y] = "A\t"
			} else if y <= x && x+y >= n-1 {
				pola[x][y] = "B\t"
			} else {
				pola[x][y] = "\t"
			}
		}
		genap -= 2
	}

	//Mencetak
	for x := 0; x < len(pola); x++ {
		for y := 0; y < len(pola[x]); y++ {
			fmt.Print(pola[x][y])
		}
		fmt.Println()
	}
}

func soal2() {
	fmt.Print("Masukan jumlah n = ")
	fmt.Scanln(&n)

	//Membuat Baris
	pola := make([][]string, n)

	//Membuat kolom
	for i := range pola {
		pola[i] = make([]string, n)
	}

	var ganjil2 int = 1
	var genap2 int = (n * 2) - 2

	//Memasukan nilai kedalam array
	for x := 0; x < len(pola); x++ {
		var ganjil1 int = 1
		var genap1 int = 0
		for y := 0; y < len(pola); y++ {
			if x == y {
				pola[x][y] = strconv.Itoa(ganjil1) + "\t"
			} else if y == n/2 {
				pola[x][y] = strconv.Itoa(ganjil2) + "\t"
			} else if x == n/2 {
				pola[x][y] = strconv.Itoa(genap1) + "\t"
			} else if x+y == n-1 {
				pola[x][y] = strconv.Itoa(genap2) + "\t"
			} else {
				pola[x][y] = "\t"
			}
			ganjil1 += 2
			genap1 += 2
		}
		ganjil2 += 2
		genap2 -= 2
	}

	//Mencetak
	for x := 0; x < len(pola); x++ {
		for y := 0; y < len(pola[x]); y++ {
			fmt.Print(pola[x][y])
		}
		fmt.Println()
	}
}

func soal3() {
	fmt.Print("Masukan jumlah n = ")
	fmt.Scanln(&n)

	//Membuat Baris
	pola := make([][]string, n)

	//Membuat kolom
	for i := range pola {
		pola[i] = make([]string, n)
	}

	//var ganjil int = 1
	//var genap int = 0

	//Memasukan nilai kedalam array
	for x := 0; x < len(pola); x++ {
		var ganjil int = 1
		for y := 0; y < len(pola); y++ {
			if x >= y { //garis miring kanan kebawah
				pola[x][y] = strconv.Itoa(ganjil) + "\t"
			} else {
				pola[x][y] = "\t"
			}
			ganjil += 2
		}
	}

	//Mencetak
	for x := 0; x < len(pola); x++ {
		for y := 0; y < len(pola[x]); y++ {
			fmt.Print(pola[x][y])
		}
		fmt.Println()
	}
}

func soal4() {
	fmt.Print("Masukan jumlah n = ")
	fmt.Scanln(&n)

	//Membuat Baris
	pola := make([][]string, n)

	//Membuat kolom
	for i := range pola {
		pola[i] = make([]string, n)
	}

	// var ganjil1 int = 1
	// var ganjil2 int = 7

	//Memasukan nilai kedalam array
	for x := 0; x < len(pola); x++ {
		var ganjil1 int = 1
		var ganjil2 int = n

		for y := 0; y < len(pola); y++ {
			if y <= n/2 { //garis miring kanan kebawah
				pola[x][y] = strconv.Itoa(ganjil1) + "\t"
				ganjil1 += 2
			} else if y >= n/2 { //garis miring kanan kebawah
				ganjil2 -= 2
				pola[x][y] = strconv.Itoa(ganjil2) + "\t"
			} else {
				pola[x][y] = "\t"
			}
		}
	}

	//Mencetak
	for x := 0; x < len(pola); x++ {
		for y := 0; y < len(pola[x]); y++ {
			fmt.Print(pola[x][y])
		}
		fmt.Println()
	}
}
