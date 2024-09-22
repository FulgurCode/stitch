package handler

import (
	"fmt"
	"net/http"

	"github.com/FulgurCode/stitch/models"
	"github.com/FulgurCode/stitch/pkg/mysql"
	"github.com/FulgurCode/stitch/utils"
	"github.com/FulgurCode/stitch/view/admin"
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
	var loggedIn = utils.GetSessionValue(c, "auth", "isLoggedIn")
	if loggedIn != nil && loggedIn.(bool) {
		return c.Redirect(http.StatusSeeOther, "/admin")
	}

	return utils.Render(c, component)
}

// Admin Login request
func AdminLoginPost(c echo.Context) error {
	var body models.Admin
	var err = c.Bind(&body)
	if err != nil {
		fmt.Println(err)
	}

	user, err := mysql.GetAdminUser(body.Username)
	var correct = utils.BcryptCompare(body.Password, user.Password)

	if !correct {
		return utils.Render(c, admin.Login())
	}

	utils.CreateSession(c, "auth")
	utils.AddSessionValue(c, "auth", "username", user.Username)
	utils.AddSessionValue(c, "auth", "isLoggedIn", true)

	return c.Redirect(http.StatusSeeOther, "/admin")
}

// Admin Logout
func AdminLogout(c echo.Context) error {
	utils.DeleteSession(c, "auth")

	return c.Redirect(http.StatusSeeOther, "/admin/login")
}

func AdminChangePassword(c echo.Context) error {
	var admin models.Admin
	var err = c.Bind(&admin)
	if err != nil {
		fmt.Println(err)
	}

	admin.Username = utils.GetSessionValue(c,"auth","username").(string)
	admin.Password= utils.BcryptGenerateHash(admin.Password)

	err = mysql.UpdateAdminPassword(admin)
	if err != nil {
		fmt.Println(err)
	}

	return c.Redirect(http.StatusSeeOther, "/admin")
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
