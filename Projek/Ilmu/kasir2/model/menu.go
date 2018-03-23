package model

import f "fmt"

type Menu struct {
	ID    int
	Nama  string
	Harga int
}

func CetakMenu(mdl Menu) {
	f.Println("Data Makanan")
	f.Println("ID\t= ", mdl.ID)
	f.Println("Nama\t= ", mdl.Nama)
	f.Println("Harga\t= ", mdl.Harga)
}

func GetDataMenu() Menu {
	mdl := Menu{
		ID:    1,
		Nama:  "Mie ayam",
		Harga: 12000,
	}

	return mdl
}
