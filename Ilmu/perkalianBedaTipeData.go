package main

import "fmt"

func main() {
	var a int = 5
	var b uint = 10
	var c float32 = 50.2687

	fmt.Println("Ini Tipe Data Int\t= ", a)
	fmt.Println("Ini Tipe Data UI\t= ", b)
	fmt.Println("Ini Tipe Data Float32\t= ", c)

	/*
		Melakukan konversi
		dengan cara :
		var <Nama Variabel> <Tipe Data> = <Tipe Data>(Nama Variabel yang berbeda tipe data)
	*/
	var d float32 = float32(a)
	//Mencetak nilai hasil konversi
	fmt.Println("Nilai d dari a\t\t= ", d)

	//perhitungan variabel konversi
	var h int = int(b * uint(c))
	fmt.Println("Hasil dari perkalian tipe data\nfloat32 dengan uint yang mempunyai nilai  10 x 50.2687 = ", h)
}
