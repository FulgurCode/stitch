package handler

import (
	"github.com/VAISHAKH-GK/cloth-shop-website/utils"
	"github.com/VAISHAKH-GK/cloth-shop-website/view"
	"github.com/labstack/echo/v4"
)

// Home handler
func Home(c echo.Context) error {
	var component = view.Index()

	return utils.Render(c, component)
}
