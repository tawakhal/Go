package models

type Selling struct {
	ID          int     `json:"ID"`
	Invoice     string  `json:"Invoice"`
	InvoiceDate string  `json:"InvoiceDate"`
	Item        int     `json:"Item"`
	Total       float64 `json:"Total"`
	Paid        float64 `json:"Paid"`
	Kembali     float64 `json:"Kembali"`
	OfficerCode string  `json:"OfficerCode"`
}
