package models

type Product struct {
	Id          string `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Category    string `json:"category" form:"category"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
}

type Stock struct {
	ProductId string `json:"productId" form:"productId"`
	S         int    `json:"S" form:"S"`
	M         int    `json:"M" form:"M"`
	L         int    `json:"L" form:"L"`
	XL        int    `json:"XL" form:"XL"`
	XXL       int    `json:"XXL" form:"XXL"`
	XXXL      int    `json:"XXXL" form:"XXXL"`
}
