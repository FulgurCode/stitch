package mysql

import (
	"fmt"

	"github.com/FulgurCode/stitch/models"
)

func MakeOrder(order models.Order) error {
	var query = fmt.Sprintf("INSERT INTO orders(id,product_id,name,address,house,pin,city,phone,payment,total) VALUES(uuid(),'%s','%s','%s','%s',%d,'%s','%s','%s',%d);", order.ProductId, order.Name, order.Address, order.House, order.Pin, order.City, order.Phone, order.Payment, order.Total)

	var _, err = Db.Exec(query)

	return err
}
