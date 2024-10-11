package mysql

import (
	"fmt"

	"github.com/FulgurCode/stitch/models"
)

// Add stock
func AddStock(stock models.Stock) error {
	var query = fmt.Sprintf("INSERT INTO stock(product_id, s,m,l,xl,xxl,xxxl,total) VALUES('%s',%d,%d,%d,%d,%d,%d,s + m + l + xl + xxl + xxxl);", stock.ProductId, stock.S, stock.M, stock.L, stock.XL, stock.XXL, stock.XXXL)

	var _, err = Db.Exec(query)

	return err
}

func UpdateStock(stock models.Stock) error {
	var query = fmt.Sprintf("UPDATE stock SET s = %d, m = %d, l = %d, xl = %d, xxl = %d, xxxl = %d, total = s + m + l + xl + xxl + xxxl WHERE product_id = '%s';", stock.S, stock.M, stock.L, stock.XL, stock.XXL, stock.XXXL, stock.ProductId)

	var _, err = Db.Exec(query)

	return err
}
