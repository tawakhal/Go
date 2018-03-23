package model

import f "fmt"

type Ruangan struct {
	ID   int
	Nama string
	Luas int
}

func CetakRuangan(r Ruangan) {
	f.Println("Data Ruangan")
	f.Println("ID\t= ", r.ID)
	f.Println("Nama\t= ", r.Nama)
	f.Println("Luas\t= ", r.Luas, " Meter")
}

func GetDataRuangan() Ruangan {
	mdl := Ruangan{
		ID:   1,
		Nama: "Xsis Academy",
		Luas: 200,
	}

	return mdl
}
