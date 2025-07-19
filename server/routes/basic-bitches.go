package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterBasicBitchRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"documentation": "hi",
		})
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "PONG HOLY CRAP LOUIS I'M ALIVE")
	})

	// Returns the server status and stuff, y'know
	e.GET("/stats", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"response":             "your mom",
			"postgres_db_response": "kinda fast",
			"redis_db_response":    "it had constipation",
		})
	})

	// Available for public instance of Flinky, may vary to someone who self-hosts Flinky
	e.GET("/export-data", func(c echo.Context) error {
		return c.String(http.StatusOK, "wip")
	})

	e.POST("/export-data/:fragment", func(c echo.Context) error {
		return c.String(http.StatusOK, "this will return some blob data and stuff")
	})

}
