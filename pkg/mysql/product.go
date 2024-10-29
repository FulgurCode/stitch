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
	DeleteStock(prodcutId)
	var query = fmt.Sprintf("DELETE FROM product WHERE id = '%s';", prodcutId)

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

func EditProduct(product models.Product) error {
	var query = fmt.Sprintf("UPDATE product SET name = '%s', category ='%s', price = %d, description = '%s' WHERE id = '%s';", product.Name, product.Category, product.Price, product.Description, product.Id)

	var _, err = Db.Exec(query)
	return err
}

func SearchProduct(search string) []models.Product {
	var query = fmt.Sprintf("SELECT id,name,category,price,description FROM product WHERE name LIKE '%%%s%%' OR description LIKE '%%%s%%' OR category LIKE '%%%s%%';", search, search, search)
	var results, err = Db.Query(query)
	if err != nil {
		fmt.Println(err)
	}

	var products []models.Product
	for results.Next() {
		var product models.Product
		results.Scan(&product.Id, &product.Name, &product.Category, &product.Price, &product.Description)
		products = append(products, product)
	}

	return products
}
