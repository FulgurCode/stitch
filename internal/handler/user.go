package handler

import (
	"fmt"
	"net/http"

	"github.com/FulgurCode/stitch/models"
	"github.com/FulgurCode/stitch/pkg/mysql"
	"github.com/FulgurCode/stitch/utils"
	"github.com/FulgurCode/stitch/view/layout"
	"github.com/FulgurCode/stitch/view/user"
	"github.com/labstack/echo/v4"
)

// Home page handler
func Home(c echo.Context) error {
	var products, err = mysql.GetProducts()
	if err != nil {
		fmt.Println(err)
	}

	var settings = mysql.GetSettings()

	var component = user.Home(products,settings)

	return utils.Render(c, component)
}

// Product page handler
func Products(c echo.Context) error {
	var products, err = mysql.GetProducts()
	if err != nil {
		fmt.Println(err)
	}

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

// Order page handler
func OrderGet(c echo.Context) error {
	var productId = c.Param("productId")
	var product, err = mysql.GetProductById(productId)
	if err != nil {
		fmt.Println(err)
	}

	var component = user.Order(product)
	return utils.Render(c, component)
}

func OrderPost(c echo.Context) error {
	var productId = c.Param("productId")
	var order models.Order
	var err = c.Bind(&order)

	product, err := mysql.GetProductById(productId)
	if err != nil {
		fmt.Println(err)

		if c.Request().Header.Get("HX-Request") == "true" {
			c.Response().Header().Set("HX-Location", "/products")
			return c.NoContent(http.StatusSeeOther)
		}
		c.Redirect(http.StatusSeeOther, "/products")
	}

	order.ProductId = productId
	order.Total = product.Price
	order.Status = "ordered"

	err = mysql.MakeOrder(order)
	if err != nil {
		fmt.Println(err)
	}

	if c.Request().Header.Get("HX-Request") == "true" {
		c.Response().Header().Set("HX-Location", fmt.Sprintf("/item/%s", productId))
		return c.NoContent(http.StatusSeeOther)
	}

	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/item/%s", productId))
}

// Search page handler
func Search(c echo.Context) error {
	var component = user.Search()

	return utils.Render(c, component)
}

// Cart page handler
func Cart(c echo.Context) error {
	var component = user.Cart()

	return utils.Render(c, component)
}

// About page handler
func About(c echo.Context) error {
	var component = user.About()

	return utils.Render(c, component)
}

// Comming Soon handler
func CommingSoon(c echo.Context) error {
	var component = layout.CommmingSoon()

	return utils.Render(c, component)
}

// Not Found handler
func NotFound(c echo.Context) error {
	var component = layout.NotFound()

	return utils.Render(c, component)
}
