package soal

import (
	drt "Olgi/etc"
	mdl "Olgi/model"
	"fmt"
	"strconv"
)

func Soal3() {
	var dataBaris mdl.Matrik
	fmt.Print("Masukan Nilai N = ")
	fmt.Scanln(&dataBaris.N)

	Prima := drt.GetPrima(dataBaris.N)

	Array1D := make([]int, dataBaris.N)
	nilai := 1
	//ga := 1
	ba := 4
	naa := 98
	bor := 0
	for x := 0; x < len(Array1D); x++ {
		if nilai == ba {
			naa += 2
			ba += 4
			Array1D[x] = naa
		} else {
			if nilai <= 2 {
				Array1D[x] = Prima[bor]
				bor++
			} else {
				Array1D[x] = Prima[bor]
				bor++
			}
		}
		nilai++

	}

	for x := range Array1D {
		fmt.Print(Array1D[x], ", ")
	}
	fmt.Println()
}
