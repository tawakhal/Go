package data

import (
	mdl "DBSellingGorilla/item/models"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type ItemRepository struct {
	DB *sql.DB
}

func GetAll(db *ItemRepository) []mdl.Item {
	// Mengecek apakah DB ada isinya
	fmt.Println(db.DB)

	// Mengambil data dari database dan dimasukan kedalam rows dan err
	rows, err := db.DB.Query("SELECT tblitemid,itemcode,itemname,buyingprice,sellingprice,itemamount,pieces FROM tblitem;")

	// Mengecek Error dari pengambilan database
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Menutup rows jika sudah digunakan diakhir eksekusi
	defer rows.Close()

	// Mengecek nilai dari rows
	fmt.Println(rows)

	item := []mdl.Item{}

	for rows.Next() {
		// Membuat variabel baru dari struc mdl.Item
		var it mdl.Item

		// Membuat variabel baru err dari scan data yg diambil
		err := rows.Scan(&it.ID, &it.ItemCode, &it.ItemName, &it.BuyPrice, &it.SellPrice, &it.ItemAmount, &it.Pieces)

		// Mengecek Error data yg diambil
		if err != nil {
			fmt.Println(err)
			return nil
		}
		// Memasukan data kedalam Slice item
		item = append(item, it)
	}

	// Mengembalikan nilai dari item
	return item
}
