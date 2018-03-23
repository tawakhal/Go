package models

type Officer struct {
	ID              int    `json:"ID"`
	OfficerCode     string `json:"OfficerCode"`
	OfficerName     string `json:"OfficerName"`
	OfficerPassword string `json:"OfficerPassword"`
	OfficerStatus   string `json:"OfficerStatus"`
}
