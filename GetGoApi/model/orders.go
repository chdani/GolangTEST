package model

type Order struct {
	Order_name       string `db:"order_name"`
	Product          string `db:"product"`
	Company_name     string `db:"company_name"`
	Customer_Name    string `db:"Customer_Name"`
	Created_at       string `db:"created_at"`
	Delivered_Amount string `db:"Delivered_Amount"`
	Total_Amount     string `db:"Total_Amount"`
}

type JsonOrder struct {
	Order_name       string `json:"order_name"`
	Product          string `json:"product"`
	Company_name     string `json:"company_name"`
	Customer_Name    string `json:"Customer_Name"`
	Created_at       string `json:"created_at"`
	Delivered_Amount string `json:"Delivered_Amount"`
	Total_Amount     string `json:"Total_Amount"`
}

type OrderRequest struct {
	FromDate string `json:"FromDate"`
	ToDate   string `json:"ToDate"`
}
