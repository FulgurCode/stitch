package handler

import (
	"fmt"
	"net/http"

	"github.com/FulgurCode/stitch/models"
	"github.com/FulgurCode/stitch/pkg/mysql"
	"github.com/FulgurCode/stitch/utils"
	"github.com/FulgurCode/stitch/view/admin"
	"github.com/google/uuid"
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

	if c.Request().Header.Get("HX-Request") == "true" {
		c.Response().Header().Set("HX-Location", "/admin")
		return c.NoContent(200)
	}

	return c.Redirect(http.StatusSeeOther, "/admin")
}

// Admin Logout
func AdminLogout(c echo.Context) error {
	utils.DeleteSession(c, "auth")

	return c.Redirect(http.StatusSeeOther, "/admin/login")
}

func AdminChangePassword(c echo.Context) error {
	var component = admin.AdminChangePassword()

	return utils.Render(c, component)
}

func AdminChangePasswordPost(c echo.Context) error {
	var body models.AdminPass
	var err = c.Bind(&body)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusSeeOther, "/admin")
	}

	body.Username = utils.GetSessionValue(c, "auth", "username").(string)

	a, err := mysql.GetAdminUser(body.Username)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusSeeOther, "/admin")
	}

	if !utils.BcryptCompare(body.OldPassword, a.Password) {
		var component = admin.AdminChangePassword()

		return utils.Render(c, component)
	}

	a.Password = utils.BcryptGenerateHash(body.NewPassword)

	err = mysql.UpdateAdminPassword(a)
	if err != nil {
		fmt.Println(err)
	}

	return c.Redirect(http.StatusSeeOther, "/admin")
}

// Admin Products Handler
func AdminProducts(c echo.Context) error {
	var products, err = mysql.GetProducts()
	if err != nil {
		var component = admin.AdminProducts([]models.Product{})

		return utils.Render(c, component)
	}

	var component = admin.AdminProducts(products)

	return utils.Render(c, component)
}

// Add Product GET
func AddProductGet(c echo.Context) error {
	var component = admin.AdminAddProduct()

	return utils.Render(c, component)
}

// Add Product POST
func AddProduct(c echo.Context) error {
	var product models.Product
	c.Bind(&product)

	product.Id = uuid.NewString()

	var err = mysql.AddProduct(product)
	if err != nil {
		fmt.Println(err)
	}

	var stock models.Stock = models.Stock{
		ProductId: product.Id,
		S:         0,
		M:         0,
		L:         0,
		XL:        0,
		XXL:       0,
		XXXL:      0,
	}

	err = mysql.AddStock(stock)
	if err != nil {
		fmt.Println(err)
	}

	if c.Request().Header.Get("HX-Request") == "true" {
		c.Response().Header().Set("HX-Location", "/admin/products")
		return c.NoContent(200)
	}

	return c.Redirect(http.StatusSeeOther, "/admin/products")
}

func AdminUpdateStockPost(c echo.Context) error {
	var stock models.Stock
	c.Bind(&stock)

	var err = mysql.UpdateStock(stock)
	if err != nil {
		fmt.Println(err)
	}

	var component = admin.AdminStock()

	return utils.Render(c, component)
}

// Admin Item Handler
func AdminItem(c echo.Context) error {
	var component = admin.AdminItem()

	return utils.Render(c, component)
}

// Admin Orders Handler
func AdminOrders(c echo.Context) error {
	var component = admin.AdminOrders()

	return utils.Render(c, component)
}

// Admin Stock Handler
func AdminStock(c echo.Context) error {
	var component = admin.AdminStock()

	return utils.Render(c, component)
}

// Admin Setting Handler
func AdminSettings(c echo.Context) error {
	var component = admin.AdminSettings()

	return utils.Render(c, component)
}
