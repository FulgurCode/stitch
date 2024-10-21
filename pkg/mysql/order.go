package mysql

import (
	"fmt"

	"github.com/FulgurCode/stitch/models"
)

func MakeOrder(order models.Order) error {
	var query = fmt.Sprintf("INSERT INTO orders(id,product_id,name,address,house,pin,city,phone,payment,quantity,total,status) VALUES(uuid(),'%s','%s','%s','%s',%d,'%s','%s','%s',%d, %d,'%s');", order.ProductId, order.Name, order.Address, order.House, order.Pin, order.City, order.Phone, order.Payment, order.Quantity, order.Total, order.Status)

	var _, err = Db.Exec(query)

	return err
}

func GetOrders() ([]models.Order, error) {
	var query = fmt.Sprintf("SELECT orders.id,product_id,orders.name,address,house,pin,city,phone,payment,quantity,total,status,product.name,product.price FROM orders LEFT JOIN product ON orders.product_id = product.id;")

	var orders []models.Order
	var result, err = Db.Query(query)

	if err != nil {
		return orders, err
	}

	for result.Next() {
		var order models.Order
		result.Scan(&order.Id, &order.ProductId, &order.Name, &order.Address, &order.House, &order.Pin, &order.City, &order.Phone, &order.Payment, &order.Quantity, &order.Total, &order.Status, &order.ProductName, &order.ProductPrice)

		orders = append(orders, order)
	}

	return orders, err
}
