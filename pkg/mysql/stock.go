package mysql

import (
	"fmt"

	"github.com/FulgurCode/stitch/models"
)

// Add stock
func AddStock(stock models.Stock) error {
	var query = fmt.Sprintf("INSERT INTO stock VALUES('%s',%d,%d,%d,%d,%d,%d,s + m + l + xl + xxl + xxxl);", stock.ProductId, stock.S, stock.M, stock.L, stock.XL, stock.XXL, stock.XXXL)

	var _, err = Db.Exec(query)

	return err
}
