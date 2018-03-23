package main

import (
	sl "Olgi/soal"
	"fmt"
)

var (
	pilih      bool = true
	jawab      bool = true
	pilihan    int
	pilih_soal int
)

func main() {
	for pilih == true {
		pilih = false
		jawab = true
		fmt.Println("Soal Mana yang mau anda Lihat ?")
		for i := 1; i <= 6; i++ {
			fmt.Printf("[%d] Soal Nomor %d \n", i, i)
		}
		fmt.Print("Jawaban Anda : ")
		fmt.Scanln(&pilih_soal)

		switch pilih_soal {
		case 1:
			{
				sl.Soal1()
			}
		case 2:
			{
				sl.Soal2()
			}
		case 3:
			{
				sl.Soal3()
			}
		case 4:
			{
				sl.Soal4()
			}
		case 5:
			{
				sl.Soal5()
			}
		case 6:
			{
				sl.Soal6()
			}
		default:
			{
				fmt.Println("[Jawaban Tidak ada]Tolong Jawab dengan yang benar !!")
			}
		}

		for jawab == true {
			jawab = false
			pilih = false
			fmt.Println("Apakah Ingin Melihat Soal lain ?")
			fmt.Println("[1] Ya")
			fmt.Println("[2] Tidak")
			fmt.Print("Jawaban Anda : ")
			fmt.Scanln(&pilihan)
			switch {
			case pilihan == 1 || pilihan == 2:
				{
					jawab = false
					pilih = false
				}
			default:
				{
					fmt.Println("[Jawaban Tidak ada]Tolong Jawab dengan yang benar !!\nJawab dengan Angka \"1\" / \"2\"")
					jawab = true
				}
			}
		}
		switch pilihan {
		case 1:
			{
				pilih = true
			}
		case 2:
			{
				pilih = false
			}
		default:
			{
				pilih = true
			}
		}
	}
}
