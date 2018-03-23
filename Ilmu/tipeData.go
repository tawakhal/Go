package main

import (
	"fmt"
)

func main() {
	/*
		Bisa menggunakan dengan "var <Nama Variabel> string"
		ataupun "var <nama Variabel>"
	*/
	var kalimat = "Olgi tawakhal"
	fmt.Printf("Tipe Data String		=  %s\n", kalimat)

	/*
		Dalam menentukan desimal kita menggunakan tipe data yaitu
		float32 / float64 dan mereka punya masing masing range dan jangkauan berbeda
	*/
	var desimal float32 = 2.0721232
	fmt.Printf("Tipe Data Desimal		=  %.2f\n", desimal)

	/*
		Dalam menentukan bilangan integer / angka
		digolang mempunyai banyak tipe data angka
		yaitu : int8, int16, int32, int64
		dengan masing masing range / jangkaunya
	*/
	var long int64 = -1283217328173812123
	fmt.Println("Tipe Data Long			= ", long)

	/*
		uint adalah tipe data yang tidak boleh minus melainkan
		harus 0 - xxxxxxx dengan nilai positif
		di uint ada beberapa tipe data
		yaitu : uint8, uint16, uint32, int64
	*/
	var bilangan uint64 = 12312321321312
	fmt.Println("Tipe Data Bilangan apa tau dah 	= ", bilangan)

	/*
		dalam menentukan tipe data boolean kita dapat melakukan
		dengan cara "var <Nama Variabel> bool" ataupun "var <Nama Variabel>"
	*/
	var boleh bool = true
	fmt.Printf("Tipe Data Boolean		=  %t\n", boleh)
	if boleh == true {
		fmt.Println("boleh kah ? Jawab 		=  Olgi Ganteng")
	}

	/*
		dalam golang kita bisa mendeklarasikan variabel dalam 1 line
	*/
	var angka1, angka2, angka3 = 0, 1, 2
	fmt.Println("angka 1 = ", angka1, "angka 2 =", angka2, "angka 3 =", angka3)
}
