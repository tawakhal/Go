package soal

import (
	drt "Olgi/etc"
	mdl "Olgi/model"
	"fmt"
	"strconv"
)

func Soal4() {
	var dataBaris mdl.Matrik
	fmt.Print("Masukan Nilai N = ")
	fmt.Scanln(&dataBaris.N)

	Array1D := make([]int, dataBaris.N)
	nilai := 1
	//ga := 1
	asd := 1
	for x := 0; x < len(Array1D); x++ {
		if nilai%2 == 1 {

			Array1D[x] = asd
			asd += 5
		} else {
			Array1D[x] = asd
			asd -= 2

		}
		nilai++

	}

	for x := range Array1D {
		fmt.Print(Array1D[x], ", ")
	}
	fmt.Println()
}
