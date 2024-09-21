package handler

import (
	"fmt"
	"net/http"

	"github.com/VAISHAKH-GK/benster-website/models"
	"github.com/VAISHAKH-GK/benster-website/pkg/mysql"
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

// Admin Login request
func AdminLoginPost(c echo.Context) error {
	var body models.Admin
	var err = c.Bind(&body)
	if err != nil {
		fmt.Println(err)
	}

	err, user := mysql.GetAdminUser(body.Username)
	var correct = utils.BcryptCompare(body.Password, user.Password)

	if correct {
		return c.Redirect(http.StatusMovedPermanently, "/admin")
	}

	return utils.Render(c, admin.Login())
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
