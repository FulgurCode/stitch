package handlers

import (
	"context"

	"github.com/VAISHAKH-GK/cloth-shop-website/view"
	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
  var component = view.Index()
  return component.Render(context.Background(), c.Response().Writer)
}
