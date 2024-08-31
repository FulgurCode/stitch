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
func Item(c echo.Context) error { // Item
	var component = user.Item()

	return utils.Render(c, component)
}

// Admin handler
func Admin(c echo.Context) error { // Admin
	var component = admin.Admin()

	return utils.Render(c, component)
}
func AdminLogin(c echo.Context) error { // Admin Login
	var component = admin.Login()

	return utils.Render(c, component)
}
func AdminProducts(c echo.Context) error { // Admin Products
	var component = admin.AdminProducts()

	return utils.Render(c, component)
}
func AdminItem(c echo.Context) error { // Admin Item
	var component = admin.AdminItem()

	return utils.Render(c, component)
}
func AdminOrders(c echo.Context) error { // Admin Item
	var component = admin.AdminOrders()

	return utils.Render(c, component)
}
