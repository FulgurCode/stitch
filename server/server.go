package server

import (
	"log"

	"github.com/VAISHAKH-GK/benster-website/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run(port string) {
	var app = echo.New()

	// Remove trailing slash to avoid 404 errors
	app.Pre(middleware.RemoveTrailingSlash())

	// add static files
	app.Static("/static", "assets")

	// echo routes
	app.GET("/", handler.Home)
	app.GET("/products", handler.Products)
	app.GET("/item", handler.Item)

	app.GET("/admin", handler.Admin)
	app.POST("/admin/login", handler.AdminLoginPost)
	app.GET("/admin/login", handler.AdminLogin)
	app.GET("/admin/products", handler.AdminProducts)
	app.GET("/admin/item", handler.AdminItem)
	app.GET("/admin/orders", handler.AdminOrders)
	app.GET("/admin/settings", handler.AdminSettings)

	app.GET("/*", handler.NotFound)

	log.Fatal(app.Start(":" + port))
}
