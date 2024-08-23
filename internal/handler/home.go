package handler

import (
	"github.com/VAISHAKH-GK/cloth-shop-website/utils"
	"github.com/VAISHAKH-GK/cloth-shop-website/view/admin"
	"github.com/VAISHAKH-GK/cloth-shop-website/view/user"
	"github.com/labstack/echo/v4"
)

// Home handler
func Home(c echo.Context) error {
	var component = user.Index()

	return utils.Render(c, component)
}

func Admin(c echo.Context) error {
	var component = admin.Index()

	return utils.Render(c, component)
}
