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
	var _, err = Db.Exec(query, productId)

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

func GetStock(productId string) (map[string]int, error) {
	var query = "SELECT s,m,l,xl,xxl,xxxl,total FROM stock WHERE product_id = ?;"

	var stock = map[string]int{}
	var result, err = Db.Query(query, productId)
	if err != nil {
		return stock,err
	}

	var s,_ = result.Columns()
	result.Next()
	var scans = make([]int, len(s))
	result.Scan(&scans[0], &scans[1], &scans[2], &scans[3], &scans[4], &scans[5], &scans[6])

	for i,v := range scans {
		stock[s[i]] = v
	}

	return stock,err
}
