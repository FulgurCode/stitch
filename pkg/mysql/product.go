package mysql

import (
	"fmt"

	"github.com/FulgurCode/stitch/models"
)

func AddProduct(product models.Product) error {
	var query = fmt.Sprintf("INSERT INTO product(id,name,category,price,description) VALUES('%s','%s','%s',%d,'%s');", product.Id, product.Name, product.Category, product.Price, product.Description)

	var _, err = Db.Exec(query)

	return err
}

func GetProducts() ([]models.Product, error) {
	var products []models.Product
	var query = fmt.Sprintf("SELECT id,name,category,price,description FROM product;")

	var result, err = Db.Query(query)
	if err != nil {
		return products, err
	}

	for result.Next() {
		var p models.Product
		result.Scan(&p.Id, &p.Name, &p.Category, &p.Price, &p.Description)

		products = append(products, p)
	}

	return products, err
}

func DeleteProduct(prodcutId string) error {
	var query = fmt.Sprintf("DELETE FROM product WHERE product_id = %s;", prodcutId)

	var _, err = Db.Exec(query)

	return err
}

func GetProductById(id string) (models.Product, error) {
	var product models.Product

	var query = fmt.Sprintf("SELECT id,name,category,price,description FROM product WHERE id = '%s';", id)
	var result, err = Db.Query(query)
	if err != nil {
		return product, err
	}

	result.Next()
	result.Scan(&product.Id, &product.Name, &product.Category, &product.Price, &product.Description)

	return product, nil
}
