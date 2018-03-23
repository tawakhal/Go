package main

import (
	mdl "P_restoku/model"
	f "fmt"
)

var (
	ID         int
	Nama       string
	Pengalaman string
)

func main() {
	f.Print("Masukan ID Kasir\t= ")
	f.Scanln(&ID)
	f.Print("Masukan Nama Kasir\t= ")
	f.Scanln(&Nama)
	f.Print("Masukan Lama Pengalaman\t= ")
	f.Scanln(&Pengalaman)

	md := mdl.GetDataKasir(ID, Nama, Pengalaman)
	mdl.Kasir.CetakKasir(md)
	dataMenu := mdl.GetDataMenu()
	mdl.CetakMenu(dataMenu)
	//mdl.CetakMenu(mdl.GetDataMenu);
	dataRuangan := mdl.GetDataRuangan()
	mdl.CetakRuangan(dataRuangan)
	//mdl.CetakRuangan(mdl.GetDataRuangan);
}
