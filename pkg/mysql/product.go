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

func GetProducts() ([]models.Product, error) {
	var products []models.Product
	var query = fmt.Sprintf("SELECT name,category,price,description FROM product;")

	var result, err = Db.Query(query)
	if err != nil {
		return products, err
	}

	for result.Next() {
		var p models.Product
		result.Scan(&p.Name, &p.Category, &p.Price, &p.Description)

		products = append(products, p)
	}

	return products, err
}