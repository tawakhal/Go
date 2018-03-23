package models

type Item struct {
	ID         int    `json:"tblitemid"`
	ItemCode   string `json:"itemcode"`
	ItemName   string `json:"itemname"`
	BuyPrice   string `json:"buyingprice"`
	SellPrice  string `json:"sellingprice"`
	ItemAmount int    `json:"itemamount"`
	Pieces     string `json:"pieces"`
}
