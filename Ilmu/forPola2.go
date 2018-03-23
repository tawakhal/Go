package main

import "fmt"

var (
	p, t, n int
)

func main() {
	// bintangbintang()
	// polaMatrik()
	// polaApadah()
	// polaInputan()
	// polahurufL()
	// polahurufKotak()
	// polaGaris()
	// polaGarisMirror()
	// polaSilang()
	// polahurufKotakdanX()
	polaTambah()
	// olgi()
}

func polaTambah() {
	fmt.Println("Membuat Pola \"TAMBAH\" sesuai Inputan")
	fmt.Print("Masukan Jumlah yang anda inginkan	= ")
	fmt.Scanln(&n)
	fmt.Println()

	var tengah = n/2 + 1

	for x := 1; x <= n; x++ {
		for y := 1; y <= n; y++ {
			if x == y {
				fmt.Print(x+y-1, "\t")
			} else if x+y == n+1 {
				fmt.Print((x-1)*2, "\t")
			} else if y == tengah {
				fmt.Print(x, "\t")
			} else if x == tengah {
				fmt.Print(y, "\t")
			} else {
				fmt.Print("\t")
			}
		}
		fmt.Println()
	}
}

func polahurufKotakdanX() {
	fmt.Println("Membuat Pola \"Kotak X\" sesuai Inputan")
	fmt.Print("Masukan Jumlah yang anda inginkan	= ")
	fmt.Scanln(&n)
	fmt.Println()

	for x := 1; x <= n; x++ {
		for y := 1; y <= n; y++ {
			if x == n || x == 1 {
				fmt.Print(" ", x, " ")
			} else if y == n || y == 1 {
				fmt.Print(" ", x, " ")
			} else if x == y {
				fmt.Print(" ", x, " ")
			} else if x+y == n+1 {
				fmt.Print(" ", x, " ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
}

func polaSilang() {
	fmt.Println("Membuat Pola \"GARIS MIRROR\" sesuai Inputan")
	fmt.Print("Masukan Jumlah yang anda inginkan	= ")
	fmt.Scanln(&n)
	fmt.Println()

	for x := 1; x <= n; x++ {
		for y := 1; y <= n; y++ {
			if x == y {
				fmt.Print(" ", x, " ")
			} else if x+y == n+1 {
				fmt.Print(" * ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
}

func polaGarisMirror() {
	fmt.Println("Membuat Pola \"GARIS MIRROR\" sesuai Inputan")
	fmt.Print("Masukan Jumlah yang anda inginkan	= ")
	fmt.Scanln(&n)
	fmt.Println()

	for x := 1; x <= n; x++ {
		for y := 1; y <= n; y++ {
			if x+y == n+1 {
				fmt.Print(" ", x, " ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
}

func polaGaris() {
	fmt.Println("Membuat Pola \"GARIS\" sesuai Inputan")
	fmt.Print("Masukan jumlah \t= ")
	fmt.Scanln(&n)
	fmt.Println()
	fmt.Println("Hasil : ")

	for x := 1; x <= n; x++ {
		for y := 1; y <= n; y++ {
			if x == y {
				fmt.Print(" ", x, " ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
}

func polahurufL() {
	fmt.Println("Membuat Pola Huruf \"L\" sesuai Inputan")
	fmt.Print("Masukan Tinggi L\t= ")
	fmt.Scanln(&t)
	fmt.Println()
	fmt.Println("Hasil : ")

	for i := 1; i <= t; i++ {
		for j := 0; j < t; j++ {
			if i == t {
				fmt.Print("L ")
			} else if i != t && j < 1 {
				fmt.Print("L  ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func polahurufKotak() {
	fmt.Println("Membuat Pola \"KOTAK\" sesuai Inputan")
	fmt.Print("Masukan Tinggi L\t= ")
	fmt.Scanln(&t)
	fmt.Println()
	fmt.Println("Hasil : ")

	for i := 1; i <= t; i++ {
		for j := 1; j <= t; j++ {
			if i == t || i == 1 {
				fmt.Print("0 ")
			} else if i != t && j < 2 {
				fmt.Print("0 ")
			} else if i != t && j == t {
				fmt.Print("0 ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func polaInputan() {
	fmt.Println("Membuat Pola Bintang sesuai Inputan")
	fmt.Print("Masukan Panjang\t= ")
	fmt.Scanln(&p)
	fmt.Print("Masukan Lebar\t= ")
	fmt.Scanln(&t)
	fmt.Println()
	fmt.Println("Hasil : ")

	for i := 0; i < t; i++ {
		for j := 0; j < t; j++ {
			fmt.Print("*  ")
		}
		fmt.Println()
	}

}

func bintangbintang() {
	//baris
	for i := 1; i <= 3; i++ {
		//kolom
		for j := 1; j <= 3; j++ {
			fmt.Print("*\t")
		}
		fmt.Println()
	}
}

func polaMatrik() {
	//baris
	for i := 1; i <= 10; i++ {
		//kolom
		for j := 1; j <= 10; j++ {
			fmt.Print(j+i, "	")
		}
		fmt.Println()
	}
}

func polaApadah() {
	//baris
	for i := 1; i <= 10; i++ {
		//kolom
		for j := 1; j <= 10; j++ {
			if j < i {
				fmt.Print("   ")
			} else {
				fmt.Print("*  ")
			}
		}
		fmt.Println()
	}
}

func olgi() {
	fmt.Println("OOOOO", "\t", "L    ")
	fmt.Println("O   O", "\t", "L    ")
	fmt.Println("O   O", "\t", "L    ")
	fmt.Println("O   O", "\t", "L    ")
	fmt.Println("OOOOO", "\t", "LLLLL")
}
