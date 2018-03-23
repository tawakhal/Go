package model

import f "fmt"

type Kasir struct {
	ID             int
	Nama           string
	LamaPengalaman string
}

func (K1 Kasir) CetakKasir() {
	f.Println("Data Kasir")
	f.Println("ID\t\t= ", K1.ID)
	f.Println("Nama\t\t= ", K1.Nama)
	f.Println("Lama Pengalaman\t= ", K1.LamaPengalaman, " Tahun")
}

func GetDataKasir(ID int, Nama string, Pengalaman string) Kasir {
	var K1 Kasir
	K1.ID = ID
	K1.Nama = Nama
	K1.LamaPengalaman = Pengalaman
	return K1
}
