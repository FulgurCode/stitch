package handler

import (
	"fmt"

	"github.com/FulgurCode/stitch/pkg/mysql"
	"github.com/FulgurCode/stitch/utils"
	"github.com/FulgurCode/stitch/view/layout"
	"github.com/FulgurCode/stitch/view/user"
	"github.com/labstack/echo/v4"
)

// Home page handler
func Home(c echo.Context) error {
	var component = user.Home()

	return utils.Render(c, component)
}

// Product page handler
func Products(c echo.Context) error {
	var products, _ = mysql.GetProducts()
	var component = user.Products(products)

	return utils.Render(c, component)
}

// Item page handler
func Item(c echo.Context) error {
	var productId = c.Param("productId")
	var product, err = mysql.GetProductById(productId)
	if err != nil {
		fmt.Println(err)
	}

	var component = user.Item(product)
	return utils.Render(c, component)
}

// Item page handler
func Order(c echo.Context) error {
	var component = user.Order()

	return utils.Render(c, component)
}

// Not Found handler
func NotFound(c echo.Context) error {
	var component = layout.NotFound()

	return utils.Render(c, component)
}
