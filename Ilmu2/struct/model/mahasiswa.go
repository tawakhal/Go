package model

import f "fmt"

type Mahasiswa struct { // Huruf Besar supaya  Public
        ID     int
        Nama   string
        Alamat string
}

// Mencetak Mahasiswa
func CetakMahasiswa(mhs Mahasiswa) {
        f.Println("ID\t= ", mhs.ID)
        f.Println("Nama\t= ", mhs.Nama)
        f.Println("Alamat\t= ", mhs.Alamat)
}

