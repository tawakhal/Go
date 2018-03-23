/*
	Program Belajar Struct Sebagai Parameter
	Sumber : Tutor Golang #19 Struct Sebagai parameter Fungsi #2
*/

package main

import mdl "hari15/model"

func main() {
	M := mdl.Mahasiswa{
		ID:     201343500588,
		Nama:   "Olgi Tawakhal",
		Alamat: "Jln. Pedurenan Depok RT 04 RW 01 Cisalak Pasar Cimanggis Depok",
	}
	mdl.CetakMahasiswa(M)
	dosen := mdl.InputDosen();
	mdl.CetakDosen(dosen);
}
