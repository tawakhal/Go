package soal

import (
	drt "Olgi/etc"
	mdl "Olgi/model"
	"fmt"
	"strconv"
)

func Soal5() {
	var dataBaris mdl.Matrik

	fmt.Print("Masukan Nilai N = ")
	fmt.Scanln(&dataBaris.NilaiLebar)
	fmt.Print("Masukan Nilai M = ")
	fmt.Scanln(&dataBaris.NilaiPanjang)
	Fibo1 := drt.GetFibo1(200)
	Fibo2 := drt.GetFibo2(100)
	Array2D := drt.GetArray2D2(dataBaris.NilaiLebar, dataBaris.NilaiPanjang)

	for x := 0; x < dataBaris.NilaiPanjang; x++ {

		a := 10
		asd := 0
		kal := 0
		b := 10
		ku := 0
		for y := 0; y < dataBaris.NilaiLebar; y++ {
			if x+1 == 1 {
				Array2D[x][y] = strconv.Itoa(Fibo2[y])
			} else if x+1 == 2 {
				nilai2 := 100
				nilai2 -= Fibo2[asd]
				Array2D[x][y] = strconv.Itoa(nilai2)
				asd++
			} else if x+1 == 3 {

				Array2D[x][y] = strconv.Itoa(a)
				if a == 1000 {
					a = 10
				} else {
					a *= 10
				}
			} else if x+1 == 4 {
				kaa := 0
				kaa = Fibo2[kal]
				nilai2 := 100
				nilai2 -= Fibo2[kal]

				c := kaa + nilai2 + b + Fibo1[ku]
				Array2D[x][y] = strconv.Itoa(c)
				if b == 1000 {
					b = 10
				} else {
					b *= 10
				}
				ku++
			}

		}

	}

	drt.Cetak2D(Array2D)
}
