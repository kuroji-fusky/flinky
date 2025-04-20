package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/kuroji-fusky/flinky/server/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func customHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("X-Made-With", "Hugs and kisses")
		return next(c)
	}
}

func main() {
	// Load env
	cwd, cwdErr := os.Getwd()
	if cwdErr != nil {
		log.Fatalf("Error getting current working directory: %v", cwdErr)
	}

	envErr := godotenv.Load(filepath.Join(cwd, "..", ".env"))
	if envErr != nil {
		log.Fatalf("Error getting .env: %v", envErr)
	}

	// customAPIKey := os.Getenv("FLINKY_SECRET_API_TOKEN")
	// customAPIKey := os.Getenv("FLINKY_SERVER_PORT")
	allowedURLOrigins := os.Getenv("FLINKY_SERVER_CORS_ALLOWED_DOMAINS")
	if allowedURLOrigins == "" {
		allowedURLOrigins = "http://localhost:3000"
	}

	// Main server
	e := echo.New()

	e.Use(
		// CORS stuff
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     strings.Split(allowedURLOrigins, ","),
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
			AllowCredentials: true,
		}),
		// Timeout stuff
		middleware.TimeoutWithConfig(middleware.TimeoutConfig{
			Timeout: 25 * time.Second,
		}),
		// Redirect trailing slash, cuz those are freaking cringe
		middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
			RedirectCode: http.StatusPermanentRedirect,
		}),
		customHeader,
	)

	routes.BasicBitchRoutes(e)

	go func() {
		if err := e.Start(":4000"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatalf("Server failed to start: %v", err)
		}
	}()

	closeDatMf := make(chan os.Signal, 1)
	signal.Notify(closeDatMf, os.Interrupt)

	<-closeDatMf

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatalf("Server forced to shut down: %v", err)
	}

	log.Println("I sleep now")
}
