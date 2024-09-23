package mysql

import (
	"fmt"

	"github.com/FulgurCode/stitch/models"
)

func AddProduct(product models.Product) error {
	var query = fmt.Sprintf("INSERT INTO product(id,name,category,price,description) VALUES(UUID(),'%s','%s',%d,'%s');", product.Name, product.Category, product.Price, product.Description)

	var _, err = Db.Exec(query)

	return err
}
