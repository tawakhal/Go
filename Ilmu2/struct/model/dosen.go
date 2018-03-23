package model

import(
	f "fmt"
)

type Dosen struct{
	ID 	int
	Nama	string
	Alamat	string
}

func CetakDosen(mdl Dosen){
	f.Println("ID \t = ",mdl.ID);
	f.Println("Nama \t = ",mdl.Nama);
	f.Println("Alamat \t = ",mdl.Alamat);
}

func InputDosen() Dosen{
	mdl := Dosen{
		ID:2,
		Nama: "Bunga",
		Alamat: "Jakarta",
	}

	return mdl;
}
