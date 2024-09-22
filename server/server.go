package server

import (
	"log"
	"os"

	"github.com/FulgurCode/stitch/internal/handler"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run(port string) {
	var app = echo.New()

	// Remove trailing slash to avoid 404 errors
	app.Pre(middleware.RemoveTrailingSlash())

	// Creating session
	var sessionSecret = os.Getenv("SESSION_SECRET")
	app.Use(session.Middleware(sessions.NewCookieStore([]byte(sessionSecret))))

	// add static files
	app.Static("/static", "assets")

	// echo routes
	app.GET("/", handler.Home)
	app.GET("/products", handler.Products)
	app.GET("/item", handler.Item)
	app.GET("/order", handler.Order)

	app.GET("/admin", handler.Admin)

	app.GET("/admin/login", handler.AdminLogin)
	app.POST("/admin/login", handler.AdminLoginPost)
	app.GET("/admin/logout", handler.AdminLogout)
	app.POST("/admin/change-password", handler.AdminChangePassword)

	app.GET("/admin/products", handler.AdminProducts)
	app.GET("/admin/item", handler.AdminItem)
	app.GET("/admin/orders", handler.AdminOrders)
	app.GET("/admin/settings", handler.AdminSettings)

	app.GET("/*", handler.NotFound)

	log.Fatal(app.Start(":" + port))
}
