package server

import (
	"log"

	"github.com/VAISHAKH-GK/cloth-shop-website/internal/handlers"
	"github.com/labstack/echo/v4"
)

func Run(port string) {
	var app = echo.New()

  app.GET("/", handlers.Home)

	log.Fatal(app.Start(":" + port))
}
