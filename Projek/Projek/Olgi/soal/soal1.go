package soal

import (
	drt "Olgi/etc"
	mdl "Olgi/model"
	"fmt"
	"strconv"
)

func Soal1() {
	var dataBaris mdl.Matrik
	fmt.Print("Masukan Nilai N = ")
	fmt.Scanln(&dataBaris.N)

	Array1D := make([]int, dataBaris.N)
	nilai := 1
	for x := 0; x < len(Array1D); x++ {
		if nilai%2 == 1 {
			Array1D[x] = (x + 1) * (x + 1)
		} else {
			Array1D[x] = -(x + 1) * (x + 1)
		}
		nilai++

	}

	for x := range Array1D {
		fmt.Print(Array1D[x], ", ")
	}
	fmt.Println()
}
