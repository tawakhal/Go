package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//kataLingkaran()
	//breakerPembagian()
	//cekTipeData()
	//funcPeriodic()
	funcClosur()
}

func kataLingkaran() {
	var (
		luas     float64
		keliling float64
	)
	luas, keliling = hitung(10)

	fmt.Println("Luas Lingkaran\t\t= ", luas)
	fmt.Println("Keliling Lingkaran\t= ", keliling)
}

func breakerPembagian() {
	var (
		a, b float32
	)

	fmt.Print("Masukan Nilai a\t= ")
	fmt.Scanln(&a)
	fmt.Print("Masukan Nilai b\t= ")
	fmt.Scanln(&b)

	breaker(a, b)

}

//Cek type data
func cekTipeData() {
	var a, b string

	fmt.Print("Masukan Nilai a\t= ")
	fmt.Scanln(&a)
	fmt.Print("Masukan Nilai b\t= ")
	fmt.Scanln(&b)

	cekTipe(a, b)
}

func funcPeriodic() {
	hitungPeriodic(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}

func funcClosur() {
	//var angka = [5]int{1, 2, 3, 4, 5}
	var total int
	//function tanpa nama yang bisa ditaro divariabel
	var hitTotal = func(angka ...int) int {
		total := 0
		for _, nilai := range angka {
			total += nilai
		}
		return total
	}
	total = hitTotal(1, 2, 3, 4, 5)
	fmt.Println("Jumlah semua = ", total)
}

/*
	Go bisa mengembalikan lebih dari 2 nilai (bisa 2,3,4,5 ..... n)
	Mengembalikan 2 nilai dengan menggunakan 2 tipe data
	rumus :
	1. func <nama func>(<namaVariabel> <tipedata>) (<namavariabel> <tipedata>,<namavariabel> <tipedata>)
	2. func <nama func>(<namaVariabel> <tipedata>) (<tipedata>,<tipedata>)
		*jika pake cara 2, maka harus dideklarasikan dlu tipe data variabelnya

	dan dalam pengembalian bisa juga hanya menggunakan :
	1. return
	2. return <nama variabel>, <nama variabel> // jika 2 yg dikembalikan
*/

/* Cara 1
func hitung(r float64) (luas float64, keliling float64) {
	luas = math.Pi * math.Pow(r, 2)
	keliling = math.Pi * 2 * r

	return luas, keliling
}
*/

// Cara 2
func hitung(r float64) (float64, float64) {
	luas := math.Pi * math.Pow(r, 2)
	keliling := math.Pi * 2 * r

	return luas, keliling
}

//Membuat breaker
func breaker(a, b float32) {
	if b == 0 {
		fmt.Printf("Error : %2.f tidak bisa dibagi %2.f \n", a, b)
		return
	}

	hasil := a / b
	fmt.Printf("hasil dari %2.f :%2.f = %2.f\n", a, b, hasil)

}

func cekTipe(a, b string) {
	/*
		Konversikan string ke int dengan cara :
		<namavariabel>, <namavariabel untuk result untuk menentukan> := strconv.ParseInt(<variabel yg ingin dikonver>, <base>,<type size>)
			*type size = adalah ukurang
			*type size = float32, float64, int32,int64.....

			*base = base 0, base 2, base 10
	*/
	angkaA, ra := strconv.ParseInt(a, 10, 64)
	angkaB, rb := strconv.ParseInt(b, 10, 64)

	//Mengecek apakah nilai ra == nil(null) dan rb == nil(null)
	if ra == nil && rb == nil {
		fmt.Printf("%d dan %d adalah tipe data \"INTEGER\"\n", angkaA, angkaB)
		if angkaB == 0 {
			fmt.Printf("Angka ke 2 tidak boleh \"%d\"\n", angkaB)
			return
		}
		hasil := angkaA / angkaB
		fmt.Printf("Pembagian antara %d / %d adalah %d\n", angkaA, angkaB, hasil)

	} else {
		fmt.Printf("\"%s\" dan \"%s\" adalah tipe data \"STRING\"\n", a, b)
		return
	}

}

/*
 	func periodic berguna untuk mendeklarasika banyak angka
	contoh :
		func hitung(a,b,c,d,f,g,h,i,j,k,l,m,n,o,p int)
	jadi, kita tinggal menggunakan
		func hitung(a ... int)

	*variabel angka itu sama tapi tak serupa dengan slice / array
*/
func hitungPeriodic(angka ...int) {
	var total int = 0
	for angka, nilai := range angka {
		fmt.Printf("Angka ke %d\t= %d\n", angka, nilai)
		total += nilai
	}
	fmt.Println("Total\t= ", total)

}
