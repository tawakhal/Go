package main

import "fmt"

var (
	huruf     = [6]string{"JE", "KE", "TI", "FOR", "TI", "EG"}
	n     int = 0
)

func main() {
	//arrayFor()
	//arrayTidakTerhingga()
	//array2D()
	//slice()
	//sliceReference()
	//sliceCopy()
	//sliceAppend()
	//arrayInput()
	//soal1()
	//soal1for()
	//soal1Array()
	//soal2Array()
	//soal3Array()
	//soal4Array()
	//soal5Array()
	//soal6Array()
	//Map()
	mapSlice()
}

func arrayFor() {
	fmt.Println("Mengeluarkan nilai array dengan for dan len(array)")
	kalimat := ""
	/*
		 len(nama_array) adalah sebuah length dari array
			contoh : var huruf [3]string
					 Maka, len(huruf) = 3
	*/
	for i := 0; i < len(huruf); i++ {
		fmt.Println("Data ke ", i, " =", huruf[i])
		kalimat += huruf[i]
	}
	fmt.Println(kalimat)
}

func arrayTidakTerhingga() {
	fmt.Println("Array Tak Terhingga")
	//Array tak terhingga yang mempunyai nilai dan tidak bisa ditambah
	var kalimat = [...]string{"halo", "ini", "array", "tidak", "terhingga"}

	fmt.Println("Array Tak Terhingga mempunyai jumlah saat ini\t=", len(kalimat))
	for i := 0; i < len(kalimat); i++ {
		fmt.Println("Dengan nilai yaitu :\nNilai [", i, "] = ", kalimat[i])
	}

	/*
		Di array Tak terhingga kita tidak bisa menambahkan seperti :
		kalimat[len(kalimat)+1] = "MANTAB"
		jika ingin, kita bisa menggunakan Slice bukan array
	*/

}

func array2D() {
	var nama = [2][2]string{{"Olgi", "Tawakhal"}, {"Agung", "Rahmatio"}}
	for x := 0; x < len(nama); x++ {
		for y := 0; y < len(nama[x]); y++ {
			fmt.Println("INDEX KE [", x, ",", y, "]\t= ", nama[x][y])
		}

	}
}

func slice() {
	// Mendeklarasikan sebuah Slice
	var pola = []string{"1", "2", "3"}

	/*
		Membuat jumlah baris disebuah Slice sesuai n
		dalam menggunakan slice kita harus menggunakan for - range
	*/

	for i, a := range pola {
		fmt.Println("INDEX KE ", i, " = ", a)
	}
}

func sliceReference() {
	//Reference
	var huruf1 = []string{"A", "B", "C", "D"}
	/*
		Mendeklarasikan Slice yang nilainya diambil dari huruf1
		Slice huruf2 mengambil format dari huruf1 dan juga index dan nilai dari index
	*/
	var huruf2 = huruf1[1:4]

	fmt.Println("SLICE huruf1 : ")
	fmt.Println(huruf1)
	fmt.Println("SLICE huruf2 : ")
	fmt.Println(huruf2)

	var huruf3 = huruf2[1:2]

	fmt.Println("SLICE huruf3 : ")
	fmt.Println(huruf3)

	huruf3[0] = "Z"
	fmt.Println("SLICE setelah diganti C menjadi Z ")
	fmt.Println(huruf1)
	fmt.Println(huruf2)
	fmt.Println(huruf3)

}

func sliceCopy() {
	var huruf1 = []string{"i", "u", "a", "G", "H"}                      // len = 4
	var huruf2 = []string{"A", "E", "ASDA", "ASDSA", "ASDSA", "ASDASD"} // len = 5

	/*
		Copy itu adalah hasil dari yang di copy
		copy(len(slice target),len(slice Asal))
		dengan sintaks :
			var <namavariabel> = copy(<slicetarget>,<sliceAsal>)
	*/
	fmt.Println(huruf1)
	fmt.Println(huruf2)

	//var huruf3 = copy(huruf1, huruf2)
	var huruf3 = copy(huruf2, huruf1)

	fmt.Println(huruf1)
	fmt.Println(huruf2)
	fmt.Println(huruf3)
	//fmt.Println(huruf4)

}

func sliceAppend() {
	var angka1 = []int{1, 2, 3, 4}
	var huruf1 = []string{"A", "B"}

	var angka2 = append(angka1, 5)
	var huruf2 = append(huruf1, "C")

	fmt.Println(angka2)
	fmt.Println(huruf2)
}

func arrayInput() {
	fmt.Println("Membuat pola menggunakan array dan slice dengan inputan keyboard")
	fmt.Print("Masukan jumlah n = ")
	fmt.Scanln(&n)

	//Membuat slice 2D
	pola := make([][]string, n) // membuat Baris dlu berdasarkan nilai n. jadi = pola[n][]

	for d := range pola { // membuat Kolom berdasarkan nilai n. jadi = pola[n][n]
		pola[d] = make([]string, n)
	}

	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			pola[x][y] = " * "
		}
	}

	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			fmt.Print(pola[x][y])
		}
		fmt.Println()
	}
}

func soal1() {
	var pola = [9][9]string{
		{"*", " ", " ", " ", " ", " ", " ", " ", " "},
		{"*", "*", " ", " ", " ", " ", " ", " ", " "},
		{"*", "*", "*", " ", " ", " ", " ", " ", " "},
		{"*", "*", "*", "*", " ", " ", " ", " ", " "},
		{"*", "*", "*", "*", "*", " ", " ", " ", " "},
		{"*", "*", "*", "*", "*", "*", " ", " ", " "},
		{"*", "*", "*", "*", "*", "*", "*", " ", " "},
		{"*", "*", "*", "*", "*", "*", "*", "*", " "},
		{"*", "*", "*", "*", "*", "*", "*", "*", "*"},
	}

	for x := 0; x < len(pola); x++ {
		for y := 0; y < len(pola[x]); y++ {
			fmt.Print("", pola[x][y])
		}
		fmt.Println()
	}
}

func soal1for() {
	var pola [9][9]string

	for x := 0; x < len(pola); x++ {
		for y := 0; y < len(pola); y++ {
			if x > y {
				pola[x][y] = " * "
			} else {
				pola[x][y] = "   "
			}
		}
	}

	for i := 0; i < len(pola); i++ {
		for j := 0; j < len(pola); j++ {
			fmt.Print(pola[i][j])
		}
		fmt.Println()
	}
}

func soal1Array() {
	fmt.Print("Masukan jumlah n = ")
	fmt.Scanln(&n)

	pola := make([][]string, n)

	for i := range pola {
		pola[i] = make([]string, n)
	}

	for x := 0; x < len(pola); x++ {
		for y := 0; y < len(pola[x]); y++ {
			if x >= y {
				pola[x][y] = " * "
			} else {
				pola[x][y] = "   "
			}
		}
	}

	for x := 0; x < len(pola); x++ {
		for y := 0; y < len(pola[x]); y++ {
			fmt.Print(pola[x][y])
		}
		fmt.Println()
	}
}

func soal2Array() {
	fmt.Print("Masukan jumlah n = ")
	fmt.Scanln(&n)

	//Membuat array2D dengan slice
	pola := make([][]string, n) // baris pola[n][]

	//membuat variabel i dengan nilai index dari pola yang mempunya range(jangkauan) yang sama
	for i := range pola { // kolom pola[n][n]
		pola[i] = make([]string, n)
	}

	for x := 0; x < len(pola); x++ {
		for y := 0; y < len(pola[x]); y++ {
			if x+y >= n-1 {
				pola[x][y] = " * "
			} else {
				pola[x][y] = "   "
			}
		}
	}

	for x := 0; x < len(pola); x++ {
		for y := 0; y < len(pola[x]); y++ {
			fmt.Print(pola[x][y])
		}
		fmt.Println()
	}
}

func soal3Array() {
	fmt.Print("Masukan jumlah n = ")
	fmt.Scanln(&n)

	//Membuat array2D dengan slice
	pola := make([][]string, n)

	for i := range pola {
		pola[i] = make([]string, n)
	}

	for x := 0; x < len(pola); x++ {
		for y := 0; y < len(pola[x]); y++ {
			if x <= y && x+y <= n-1 {
				pola[x][y] = " * "
			} else if x >= y && x+y >= n-1 {
				pola[x][y] = " * "
			} else {
				pola[x][y] = "   "
			}
		}
	}

	for x := 0; x < len(pola); x++ {
		for y := 0; y < len(pola[x]); y++ {
			fmt.Print(pola[x][y])
		}
		fmt.Println()
	}
}

func soal4Array() {
	fmt.Print("Masukan jumlah n = ")
	fmt.Scanln(&n)

	//Membuat array2D dengan slice
	pola := make([][]string, n)

	for i := range pola {
		pola[i] = make([]string, n)
	}

	for x := 0; x < len(pola); x++ {
		for y := 0; y < len(pola[x]); y++ {
			if x >= y && x+y <= n-1 {
				pola[x][y] = " * "
			} else if x <= y && x+y >= n-1 {
				pola[x][y] = " * "
			} else {
				pola[x][y] = "   "
			}
		}
	}

	for x := 0; x < len(pola); x++ {
		for y := 0; y < len(pola[x]); y++ {
			fmt.Print(pola[x][y])
		}
		fmt.Println()
	}
}

func soal5Array() {
	fmt.Print("Masukan jumlah n = ")
	fmt.Scanln(&n)

	//Membuat array2D dengan slice
	pola := make([][]string, n)

	for i := range pola {
		pola[i] = make([]string, n)
	}

	for x := 0; x < len(pola); x++ {
		for y := 0; y < len(pola[x]); y++ {
			if x == y || x+y == n-1 {
				pola[x][y] = " * "
			} else if x == n-1 || y == n-1 {
				pola[x][y] = " * "
			} else if x == 0 || y == 0 {
				pola[x][y] = " * "
			} else {
				pola[x][y] = "   "
			}
		}
	}

	for x := 0; x < len(pola); x++ {
		for y := 0; y < len(pola[x]); y++ {
			fmt.Print(pola[x][y])
		}
		fmt.Println()
	}

}

func soal6Array() {
	fmt.Print("Masukan jumlah n = ")
	fmt.Scanln(&n)

	//Membuat array2D dengan slice
	pola := make([][]string, n)

	for i := range pola {
		pola[i] = make([]string, n)
	}

	for x := 0; x < len(pola); x++ {
		for y := 0; y < len(pola[x]); y++ {
			if x >= y && x <= (n-1)/2 {
				pola[x][y] = " * "
			} else if x <= y && x >= (n-1)/2 {
				pola[x][y] = " * "
			} else {
				pola[x][y] = "   "
			}
		}
	}

	for x := 0; x < len(pola); x++ {
		for y := 0; y < len(pola[x]); y++ {
			fmt.Print(pola[x][y])
		}
		fmt.Println()
	}
}

func Map() {
	/*
		Mendeklarasikan map dengan tipe data key string dan value int
		rumus : var <nama variabel> = map[<tipe data key>]<tipe data value>{}
		bisa juga dengan cara :
		1. var hariBulan = map[string]int{"januari": 31, "febuari": 28}
		2. var hariBulan = make(map[string]int)
		3. var hariBulan = *new(map[string]int)
	*/
	var hariBulan = map[string]int{"januari": 31, "febuari": 28}

	fmt.Println("Januari ada ", hariBulan["januari"], "Hari")
	fmt.Println("Febuari ada ", hariBulan["febuari"], "Hari")

	var hariJam = map[string]int{"senin": 8, "jumat": 5}

	for key, value := range hariJam {
		fmt.Println(key, "\t= ", value, " Jam")
	}
}

func mapSlice() {
	var mapslice = []map[string]string{
		{"nama": "joni", "umur": "21", "alamat": "jakarta"},
		{"nama": "olgi", "umur": "21", "alamat": "depok"},
	}

	for _, values := range mapslice {
		fmt.Println(values["nama"], "\t= ", values["alamat"])
	}
	fmt.Println("\n")
	for key, values := range mapslice {
		fmt.Println("DATA KE ", key)
		fmt.Println(values["nama"], "\t= ", values["alamat"])
	}
}
