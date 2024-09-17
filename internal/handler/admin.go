package handler

import (
	"github.com/VAISHAKH-GK/benster-website/utils"
	"github.com/VAISHAKH-GK/benster-website/view/admin"
	"github.com/labstack/echo/v4"
)

// Admin handler
func Admin(c echo.Context) error {
	var component = admin.Admin()

	return utils.Render(c, component)
}

// Admin Login Handler
func AdminLogin(c echo.Context) error {
	var component = admin.Login()

	return utils.Render(c, component)
}

// Admin Products Handler
func AdminProducts(c echo.Context) error {
	var component = admin.AdminProducts()

	return utils.Render(c, component)
}

// Admin Item Handler
func AdminItem(c echo.Context) error {
	var component = admin.AdminItem()

	return utils.Render(c, component)
}

// Admin Item Handler
func AdminOrders(c echo.Context) error {
	var component = admin.AdminOrders()

	return utils.Render(c, component)
}

// Admin Setting Handler
func AdminSettings(c echo.Context) error {
	var component = admin.AdminSettings()

	return utils.Render(c, component)
}
