package data

import (
	mdl "DBSellingGorilla/officer/models"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type OfficerRepository struct {
	DB *sql.DB
}

func GetAll(db *OfficerRepository) []mdl.Officer {
	// Mengecek apakah DB ada isinya
	fmt.Println(db.DB)

	// Mengambil data dari database dan dimasukan kedalam rows dan err
	rows, err := db.DB.Query("SELECT tblofficerid,officercode,officername,officerpassword,officerstatus FROM tblofficer;")

	// Mengecek Error dari pengambilan database
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Menutup rows jika sudah digunakan diakhir eksekusi
	defer rows.Close()

	// Mengecek nilai dari rows
	fmt.Println(rows)

	officer := []mdl.Officer{}

	for rows.Next() {
		// Membuat variabel baru dari struc mdl.Officer
		var of mdl.Officer

		// Membuat variabel baru err dari scan data yg diambil
		err := rows.Scan(&of.ID, &of.OfficerCode, &of.OfficerName, &of.OfficerPassword, &of.OfficerStatus)

		// Mengecek Error data yg diambil
		if err != nil {
			fmt.Println(err)
			return nil
		}
		// Memasukan data kedalam Slice officer
		officer = append(officer, of)
	}

	// Mengembalikan nilai dari officer
	return officer
}
