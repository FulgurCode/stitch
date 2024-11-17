package mysql

import (
	"fmt"

	"github.com/FulgurCode/stitch/models"
)

// Add stock
func AddStock(stock models.Stock) error {
	var query = "INSERT INTO stock(product_id, s,m,l,xl,xxl,xxxl,total) VALUES(?,?,?,?,?,?,?,s + m + l + xl + xxl + xxxl);"
	var _, err = Db.Exec(query, stock.ProductId, stock.S, stock.M, stock.L, stock.XL, stock.XXL, stock.XXXL)

	return err
}

func DeleteStock(productId string) error {
	var query = "DELETE FROM stock WHERE product_id = ?;"
	var _, err = Db.Exec(query,productId)

	return err
}

func UpdateStock(stock models.Stock) error {
	var query = "UPDATE stock SET s = ?, m = ?, l = ?, xl = ?, xxl = ?, xxxl = ?, total = s + m + l + xl + xxl + xxxl WHERE product_id = ?;"
	var _, err = Db.Exec(query, stock.S, stock.M, stock.L, stock.XL, stock.XXL, stock.XXXL, stock.ProductId)

	return err
}

func GetStocks() ([]models.Stock, error) {
	var query = fmt.Sprintf("SELECT product_id,s,m,l,xl,xxl,xxxl,total,product.name,product.price FROM stock LEFT JOIN product ON stock.product_id = product.id;")

	var stocks []models.Stock
	var result, err = Db.Query(query)

	if err != nil {
		return stocks, err
	}

	for result.Next() {
		var stock models.Stock
		result.Scan(&stock.ProductId, &stock.S, &stock.M, &stock.L, &stock.XL, &stock.XXL, &stock.XXXL, &stock.Total, &stock.ProductName, &stock.ProductPrice)

		stocks = append(stocks, stock)
	}

	return stocks, err
}
