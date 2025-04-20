package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterBasicBitchRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "hi",
		})
	})

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "I'm alive")
	})
}
