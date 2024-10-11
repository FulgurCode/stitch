package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CheckLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if isLoggedIn := GetSessionValue(c, "auth", "isLoggedIn"); isLoggedIn == nil || !isLoggedIn.(bool) {
			if c.Request().Header.Get("HX-Request") == "true" {
				c.Response().Header().Set("HX-Location", "/admin/login")
				return c.NoContent(200)
			}

			return c.Redirect(http.StatusSeeOther, "/admin/login")
		}

		return next(c)
	}
}
