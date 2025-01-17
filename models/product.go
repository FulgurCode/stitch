package models

type Product struct {
	Id          string `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Category    string `json:"category" form:"category"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
	Size        string `json:"size,omitempty" form:"size,omitempty"`
}

type Stock struct {
	ProductId    string `json:"productId" form:"productId"`
	ProductName  string `json:"productName" form:"productName"`
	ProductPrice int    `json:"productPrice" form:"productPrice"`
	S            int    `json:"S" form:"S"`
	M            int    `json:"M" form:"M"`
	L            int    `json:"L" form:"L"`
	XL           int    `json:"XL" form:"XL"`
	XXL          int    `json:"XXL" form:"XXL"`
	XXXL         int    `json:"XXXL" form:"XXXL"`
	Total        int    `json:"total" form:"total"`
}

type Order struct {
	Id           string `json:"id" form:"id"`
	ProductId    string `json:"productId" form:"productId"`
	ProductName  string `json:"productName" form:"productName"`
	ProductPrice int    `json:"productPrice" form:"productPrice"`
	Name         string `json:"name" form:"name"`
	Address      string `json:"address" form:"address"`
	House        string `json:"house" form:"house"`
	Pin          int    `json:"pin" form:"pin"`
	City         string `json:"city" form:"city"`
	Phone        string `json:"phone" form:"phone"`
	Size         string `json:"size" form:"size"`
	Payment      string `json:"payment" form:"payment"`
	Quantity     int    `json:"quantity" form:"quantity"`
	Total        int    `json:"total" form:"total"`
	Status       string `json:"status" form:"status"`
}
