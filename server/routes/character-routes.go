package routes

import (
	"net/http"
	// "strconv"

	"github.com/labstack/echo/v4"
)

type PostCharacterStuff struct {
	CharacterSlug string         `json:"slug" form:"slug"`
	Metadata      map[string]any `json:"metadata" form:"metadata"`
}

func RegisterCharacterRoutes(e *echo.Echo) {
	e.GET("/character/random", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"character": "random haha",
		})
	})

	// Returns the recent additions
	e.GET("/character", func(c echo.Context) error {
		// maxResults, maxResultsErr := strconv.Atoi(c.QueryParam("max_results"))
		// if maxResultsErr != nil {
		// 	maxResults = -1
		// }

		// TODO: more filter stuff for species, country, origin, holders, and stuff I'm too lazy too add rn

		return c.JSON(http.StatusOK, map[string]any{
			"data": []any{
				map[string]any{
					"message": "lol",
				},
			},
		})
	})

	// Returns whatever slug you provide lmao
	e.GET("/character/:slug", func(c echo.Context) error {
		sluggy := c.Param("slug")

		return c.JSON(http.StatusOK, map[string]string{
			"sluggy": sluggy,
		})
	})

	e.POST("/character/:slug", func(c echo.Context) (err error) {
		// sluggy := c.Param("slug")
		characterForm := new(PostCharacterStuff)

		if err := c.Bind(characterForm); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "something effed up lmao",
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	// som edits
	// e.PATCH("/character/:slug", func(c echo.Context) (err error) {
	// 	sluggy := c.Param("slug")
	// })

	// for yeeting them out
	// e.DELETE("/character/:slug", func(c echo.Context) (err error) {
	// 	sluggy := c.Param("slug")
	// })
}
