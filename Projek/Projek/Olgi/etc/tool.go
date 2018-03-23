package etc

import "fmt"

func GetArray2D(N int) [][]string {
	// Membuat Array2D
	Array2D := make([][]string, N)
	for i := range Array2D {
		Array2D[i] = make([]string, N)
	}
	return Array2D
}

func GetArray2D2(NL, NP int) [][]string {
	// Membuat Array2D
	Array2D := make([][]string, NP)
	for i := range Array2D {
		Array2D[i] = make([]string, NL)
	}
	return Array2D
}

func Cetak2D(Array2D [][]string) {
	for x := 0; x < len(Array2D); x++ {
		for y := 0; y < len(Array2D[x]); y++ {
			fmt.Print(Array2D[x][y], "\t")
		}
		fmt.Println()
	}
}
