package handler

import (
	"github.com/VAISHAKH-GK/cloth-shop-website/utils"
	"github.com/VAISHAKH-GK/cloth-shop-website/view/admin"
	"github.com/VAISHAKH-GK/cloth-shop-website/view/user"
	"github.com/labstack/echo/v4"
)

// User handler
func Home(c echo.Context) error { // Home
	var component = user.Home()

	return utils.Render(c, component)
}
func Products(c echo.Context) error { // Products
	var component = user.Products()

	return utils.Render(c, component)
}
func Item(c echo.Context) error { // Products
	var component = user.Item()

	return utils.Render(c, component)
}

// Admin handler
func Admin(c echo.Context) error {
	var component = admin.Index()

	return utils.Render(c, component)
}
