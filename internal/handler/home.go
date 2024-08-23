package handler

import (
	"github.com/VAISHAKH-GK/cloth-shop-website/utils"
	"github.com/VAISHAKH-GK/cloth-shop-website/view/admin"
	"github.com/VAISHAKH-GK/cloth-shop-website/view/user"
	"github.com/labstack/echo/v4"
)

// Home handler
func Home(c echo.Context) error {
	var component = user.Home()

	return utils.Render(c, component)
}

// Products handler
func Products(c echo.Context) error {
	var component = user.Products()

	return utils.Render(c, component)
}

// Admin handler
func Admin(c echo.Context) error {
	var component = admin.Index()

	return utils.Render(c, component)
}
