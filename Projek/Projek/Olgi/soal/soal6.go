package soal

import (
	drt "Olgi/etc"
	mdl "Olgi/model"
	"fmt"
	"strconv"
)

func Soal6() {
	var dataBaris mdl.Matrik
	fmt.Print("Masukan Nilai N = ")
	fmt.Scanln(&dataBaris.N)

	for i := 2; i <= 100; i++ {

		isPrime := true

		for j := 2; j < i; j++ {

			if i%j == 0 {

				isPrime = false
			}
		}

		if isPrime == true {

			fmt.Println(i)
		}
	}
}
