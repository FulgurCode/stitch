package server

import (
	"log"

	"github.com/VAISHAKH-GK/cloth-shop-website/internal/handler"
	"github.com/labstack/echo/v4"
)

func Run(port string) {
	var app = echo.New()

	// echo routes
	app.GET("/", handler.Home)

	app.GET("/admin", handler.Admin)

	log.Fatal(app.Start(":" + port))
}
