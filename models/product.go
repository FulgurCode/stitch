package models

type Product struct {
	Id          string `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Category    string `json:"category" form:"category"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
}
