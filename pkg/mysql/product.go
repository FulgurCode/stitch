package mysql

import (
	"fmt"

	"github.com/FulgurCode/stitch/models"
)

func AddProduct(product models.Product) error {
	var query = "INSERT INTO product(id,name,category,price,description) VALUES(?,?,?,?,?);"
	var _, err = Db.Exec(query, product.Id, product.Name, product.Category, product.Price, product.Description)

	return err
}

func GetProducts() ([]models.Product, error) {
	var products []models.Product
	var query = "SELECT id,name,category,price,description FROM product;"

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
	var query = "DELETE FROM product WHERE id = ?;"

	var _, err = Db.Exec(query, prodcutId)

	return err
}

func GetProductById(id string) (models.Product, error) {
	var product models.Product

	var query = "SELECT id,name,category,price,description FROM product WHERE id = ?;"
	var result, err = Db.Query(query, id)
	if err != nil {
		return product, err
	}

	result.Next()
	result.Scan(&product.Id, &product.Name, &product.Category, &product.Price, &product.Description)

	return product, nil
}

func EditProduct(product models.Product) error {
	var query = "UPDATE product SET name = ?, category = ?, price = ?, description = ? WHERE id = ?;"
	var _, err = Db.Exec(query, product.Name, product.Category, product.Price, product.Description, product.Id)
	return err
}

func SearchProduct(search string) []models.Product {
	var searchTerm = "%" + search + "%"
	var query = "SELECT id,name,category,price,description FROM product WHERE name LIKE ? OR description LIKE ? OR category LIKE ?;"
	var results, err = Db.Query(query, searchTerm, searchTerm, searchTerm)
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
