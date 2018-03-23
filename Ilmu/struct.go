package main

import (
	"fmt"
	"strconv"
)

/*
	LEBIH BAIK MENDEKLARASIKA SEBUAH STRUCT DILUAR FUNC / FUNC MAIN
	LEBIH BAIK DALAM HAL MAINTENANCE DAN LAINNYA.
	SEPERTI, JAVA GETTER & SETTER / POJO
*/

func main() {
	//struct1()
	//structArray()
	//structInisialisasi()
	//structPointer()
	//structDIstruct()
	//structDIstructSAMA()
	//CaraLainNgisiSubstruck()
	//AnonymousStruct()
}

func struct1() {
	//Deklarasikan Struct
	type siswa struct {
		nama string
		umur int
	}

	//cara 1
	var s1 siswa //Membuat variabel s1 yang mempunyai poperty dari struck siswa
	//cara 2 = var s1 = siswa{}
	s1.nama = "Olgi Tawakhal"
	s1.umur = 21

	fmt.Println("Nama\t= ", s1.nama)
	fmt.Println("Umur\t= ", s1.umur)
}

func structArray() {
	//Mendeklarasikan struct dengan nama "Orang"
	type orang struct {
		nama [10]string // Membuat array dengan tipedata string yang mempunyai batas 10
	}

	//Membuat variabel O dengan struct
	var O orang

	// Memasukan data kedalam sebuah object O yang mempunyai poperti array nama
	for i := 0; i < 10; i++ {
		O.nama[i] = "olgi yang ke " + strconv.Itoa(i+1)
	}

	for key, nilai := range O.nama {
		fmt.Println("Data ke ", key, "\t = ", nilai)
	}
}

func structInisialisasi() {
	//Deklarasikan Struct
	type Xsis struct {
		nama string
		umur int
	}

	//membuat variabel dari object struct
	var s1 Xsis
	//Mengisi sebuah struct s1
	s1.nama = "Olgi Tawakhal"
	s1.umur = 21

	//kita juga bisa melakukannya dengan cara seperti ini
	var s2 = Xsis{"Agung", 2}

	// atau seperti ini
	var s3 = Xsis{nama: "Zacky", umur: 25}

	fmt.Println("Peserta Bootcamp Golang !!")
	fmt.Println("STRUCT KE 1(s1)")
	fmt.Println("Nama\t= ", s1.nama)
	fmt.Println("Umur\t= ", s1.umur)
	fmt.Println("STRUCT KE 2(s2)")
	fmt.Println("Nama\t= ", s2.nama)
	fmt.Println("Umur\t= ", s2.umur)
	fmt.Println("STRUCT KE 3(s3)")
	fmt.Println("Nama\t= ", s3.nama)
	fmt.Println("Umur\t= ", s3.umur)
}

func structPointer() {
	type mahasiswa struct {
		nama string
		npm  int
	}

	var s1 = mahasiswa{"Olgi Tawakhal", 201343500}

	//Membuat struct pointer dari struct ke1 yaitu s1
	var s2 *mahasiswa = &s1

	fmt.Println("SEBELUM DIRUBAH")
	fmt.Println("Nama dari struct s1\t = ", s1.nama)
	fmt.Println("NPM dari struct s1\t = ", s1.nama)

	fmt.Println("Nama dari struct pointer s2\t = ", s2.nama)
	fmt.Println("NPM dari struct pointer s2\t = ", s2.nama)

	// Merubah nilai struct S1
	s1.nama = "Erlan Ramadhan"
	s1.npm = 201343213
	fmt.Println("\n\nNILAI S1 SETELAH DIRUBAH")
	fmt.Println("Nama dari struct s1\t = ", s1.nama)
	fmt.Println("NPM dari struct s1\t = ", s1.nama)

	fmt.Println("Nama dari struct pointer s2\t = ", s2.nama)
	fmt.Println("NPM dari struct pointer s2\t = ", s2.nama)
}

func structDIstruct() {
	type user struct {
		id_user  int
		username string
	}

	type pelayan struct {
		id_pelayan int
		nama       string
		user
	}

	//Membuat variabel struct didalam struct
	var ss1 = pelayan{}
	ss1.username = "olgi"
	ss1.id_user = 1
	ss1.id_pelayan = 912312
	ss1.nama = "Olgi Tawakhal"

	//Memanggil data menggunakan turunannya dengan inisialnya ss1.<nama turunannya>.id_user
	fmt.Println("ID Username\t= ", ss1.user.id_user)
	fmt.Println("Username\t= ", ss1.user.username)
	fmt.Println("ID Pelayan\t= ", ss1.id_pelayan)
	fmt.Println("Nama\t\t= ", ss1.nama)
	fmt.Println("Memanggil struct user tanpa memanggil user TERNYATA JUGA BISA")
	fmt.Println("ID Username\t= ", ss1.id_user)
	fmt.Println("Username\t= ", ss1.username)
}

func structDIstructSAMA() {
	type Barang1 struct {
		kode string
		nama string
		umur int
	}

	type Barang2 struct {
		Barang1
		kode string
		nama string
	}

	var ss1 = Barang2{}
	ss1.kode = "KD-12301"
	ss1.nama = "Martabak"
	ss1.umur = 22

	fmt.Println("Barang 2(UNIVERSAL/GLOBAL)")
	fmt.Println("Kode\t= ", ss1.kode)
	fmt.Println("Nama\t= ", ss1.nama)
	fmt.Println("Umur\t= ", ss1.umur)
	fmt.Println("Barang 1")
	fmt.Println("Kode\t= ", ss1.Barang1.kode)
	fmt.Println("Nama\t= ", ss1.Barang1.nama)
	fmt.Println("Umur\t= ", ss1.Barang1.umur)

	ss1.Barang1.kode = "KD-91232"
	ss1.Barang1.nama = "Telur"

	fmt.Println("BARANG 1\nSETELAH DITAMBAH MENGGUNAKAN SPESIFIK")
	fmt.Println("Kode\t= ", ss1.Barang1.kode)
	fmt.Println("Nama\t= ", ss1.Barang1.nama)
	fmt.Println("Umur\t= ", ss1.Barang1.umur)
}

func CaraLainNgisiSubstruck() {
	type orangAmerika struct {
		nama string
		asal string
	}

	type orangIndonesia struct {
		orangAmerika
		nama   string
		asal   string
		kampus string
	}

	//Membuat variabel baru dari struct
	var ss1 = orangIndonesia{}

	//cara mengisi sub struct (orangAmerika)
	//cara 1
	ss1.orangAmerika.nama = "George"

	//cara 2
	var ss2 = orangAmerika{nama: "Goerge", asal: "Amerika"}

	fmt.Println(" STRUC ORANG AMERIKA")
	fmt.Println("Nama\t: ", ss2.nama)
	fmt.Println("Asal\t: ", ss2.asal)

	//cara 2 mengisi dari struct parents
	var ss3 = orangIndonesia{orangAmerika: ss2, kampus: "UNINDRA"}
	/* ^^^ Mengisi dengan cara itu harus tidak ada variabel yang sama
	contoh :
	type orang struct{		type orang2 struct{
		nama string				orang
		asal string				nama string
	}							asal string
							}
	jika seperti contoh diatas maka jika kita panggil struct 2 nama (ss1.nama)
	maka nilainya kosong. karena lebih melihat "nama" dikelas orang2 dlu
	sebelum melihat "nama" struct orang.

	jika ingin maka harus dengan ss1.orang.nama
	*/
	fmt.Println(" STRUC ORANG INDONESIA YANG PUNYA DATA PARENTS ORANG AMERIKA")
	fmt.Println("Nama\t: ", ss3.nama)
	fmt.Println("Asal\t: ", ss3.asal)
	fmt.Println("Kampus\t: ", ss3.kampus)
	fmt.Println(" STRUC ORANG INDONESIA YANG PUNYA DATA PARENTS ORANG AMERIKA")
	fmt.Println(" DENGAN MEMANGGIL LEBIH SPESIFIK")
	fmt.Println("Nama\t: ", ss3.orangAmerika.nama)
	fmt.Println("Asal\t: ", ss3.orangAmerika.asal)
}

//anonyousstruc hanya digunakan jika dibutuhkan saja dan hanya dipakai sekali
func AnonymousStruct() {
	//Biasanya dideklarasikan diluar ataupun didalam func
	type orang struct {
		nama string
	}

	// struct ini nyaris sama seperti variabel dan tidak mempunya nama struct
	// hanya digunakan didalam func saja
	var As1 = struct {
		orang
		negara string
	}{}

	As1.nama = "olgi"
	As1.negara = "indonesia"

	fmt.Println("Nama\t = ", As1.nama)
	fmt.Println("Nama\t = ", As1.negara)
}

func sliceDANstruct() {

}
