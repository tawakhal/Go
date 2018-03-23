package main

import "fmt"

var (
	n int
)

func main() {
	fmt.Println("SOAL NO 1")
	polaSegitiga()
	fmt.Println("SOAL NO 2")
	polaSegitigaMirror()
	fmt.Println("SOAL NO 3")
	PolaJamPasir1()
	fmt.Println("SOAL NO 4")
	PolaJamPasir2()
	fmt.Println("SOAL NO 5")
	polaKotakSilang()
	fmt.Println("SOAL NO 6")
	PolaKincir()
}

func polaSegitiga() {
	fmt.Print("Masukan jumlah \t= ")
	fmt.Scanln(&n)
	fmt.Println()
	fmt.Println("Hasil : ")

	for x := 1; x <= n; x++ {
		for y := 1; y <= n; y++ {
			if y < x {
				fmt.Print(" * ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
}

func polaSegitigaMirror() {
	fmt.Print("Masukan jumlah \t= ")
	fmt.Scanln(&n)
	fmt.Println()
	fmt.Println("Hasil : ")

	for x := 1; x <= n; x++ {
		for y := 1; y <= n; y++ {
			if x+y > n+1 {
				fmt.Print(" * ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}

}

func PolaJamPasir1() {
	fmt.Print("Masukan jumlah \t= ")
	fmt.Scanln(&n)
	fmt.Println()
	fmt.Println("Hasil : ")

	for x := 1; x <= n; x++ {
		for y := 1; y <= n; y++ {
			if x+y == n || x == y {
				fmt.Print(" * ")
			} else if x+y < n+1 && x < y {
				fmt.Print(" * ")
			} else if x+y > n+1 && x > y {
				fmt.Print(" * ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
}

func PolaJamPasir2() {
	fmt.Print("Masukan jumlah \t= ")
	fmt.Scanln(&n)
	fmt.Println()
	fmt.Println("Hasil : ")

	for x := 1; x <= n; x++ {
		for y := 1; y <= n; y++ {
			if x+y == n || x == y {
				fmt.Print(" * ")
			} else if x+y < n+1 && x < y {
				fmt.Print(" * ")
			} else if x+y > n+1 && x > y {
				fmt.Print(" * ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
}

func PolaKincir() {
	fmt.Print("Masukan jumlah \t= ")
	fmt.Scanln(&n)
	fmt.Println()
	fmt.Println("Hasil : ")

	for x := 1; x <= n; x++ {
		for y := 1; y <= n; y++ {
			if y <= (n/2)+1 && x <= (n/2)+1 && y <= x {
				fmt.Print(" * ")
			} else if x >= (n/2)+1 && y >= (n/2)+1 && x <= y {
				fmt.Print(" * ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
}

func polaKotakSilang() {
	fmt.Print("Masukan jumlah \t= ")
	fmt.Scanln(&n)
	fmt.Println()
	fmt.Println("Hasil : ")

	for x := 1; x <= n; x++ {
		for y := 1; y <= n; y++ {
			if x == n || x == 1 {
				fmt.Print(" * ")
			} else if y == n || y == 1 {
				fmt.Print(" * ")
			} else if x == y {
				fmt.Print(" * ")
			} else if x+y == n+1 {
				fmt.Print(" * ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
}
