package data

import (
	mdl "DBSellingGorilla/selling/models"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type SellingRepository struct {
	DB *sql.DB
}

func GetAll(db *SellingRepository) []mdl.Selling {
	// Mengecek apakah DB ada isinya
	fmt.Println(db.DB)

	// Mengambil data dari database dan dimasukan kedalam rows dan err
	rows, err := db.DB.Query("SELECT tblsellingid,invoice,invoicedate,item,total,paid,kembali,officercode FROM tblselling;")

	// Mengecek Error dari pengambilan database
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Menutup rows jika sudah digunakan diakhir eksekusi
	defer rows.Close()

	// Mengecek nilai dari rows
	fmt.Println(rows)

	selling := []mdl.Selling{}

	for rows.Next() {
		// Membuat variabel baru dari struc mdl.Selling
		var se mdl.Selling

		// Membuat variabel baru err dari scan data yg diambil
		err := rows.Scan(&se.ID, &se.Invoice, &se.InvoiceDate, &se.Item, &se.Total, &se.Paid, &se.Kembali, &se.OfficerCode)

		// Mengecek Error data yg diambil
		if err != nil {
			fmt.Println(err)
			return nil
		}
		// Memasukan data kedalam Slice Selling
		selling = append(selling, se)
	}

	// Mengembalikan nilai dari Selling
	return selling
}
