package mysql

import (
	"github.com/FulgurCode/stitch/models"
)

func MakeOrder(order models.Order) error {
	var query = "INSERT INTO orders(id,product_id,name,address,house,pin,city,phone,payment,quantity,total,status) VALUES(uuid(),?,?,?,?,?,?,?,?,?,?,?);"
	var _, err = Db.Exec(query, order.ProductId, order.Name, order.Address, order.House, order.Pin, order.City, order.Phone, order.Payment, order.Quantity, order.Total, order.Status)

	return err
}

func GetOrdersWithStatus(status string) ([]models.Order, error) {
	var query = "SELECT orders.id,product_id,orders.name,address,house,pin,city,phone,payment,quantity,total,status,product.name,product.price FROM orders LEFT JOIN product ON orders.product_id = product.id WHERE orders.status = ?;"

	var orders []models.Order
	var result, err = Db.Query(query, status)

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

func ChangeOrderStatus(id string, status string) error {
	var query = "UPDATE orders SET status = ? WHERE id = ?;"
	var _, err = Db.Exec(query, status, id)
	return err
}

func DeleteOrder(id string) error {
	var query = "DELETE FROM orders WHERE id = ?;"
	var _, err = Db.Exec(query, id)
	return err
}
