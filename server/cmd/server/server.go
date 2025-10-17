package main

import (
	"context"
	"fmt"
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

func getEnvWithFallback[T any](envKey string, fallback T) T {
	key := os.Getenv(envKey)

	if key == "" {
		return fallback
	}

	switch any(fallback).(type) {
	case string:
		return any(key).(T)
	default:
		return fallback
	}
}

func customHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("X-Made-With", "Hugs and kisses")
		return next(c)
	}
}

func main() {
	// Load env
	cwd, _ := os.Getwd()

	envErr := godotenv.Load(filepath.Join(cwd, "..", ".env"))
	if envErr != nil {
		log.Fatalf("Error getting .env: %v", envErr)
	}

	customAPIToken := os.Getenv("FLINKY_SECRET_API_TOKEN")
	serverPort := getEnvWithFallback("FLINKY_SERVER_PORT", 6969)
	allowedURLOrigins := getEnvWithFallback("FLINKY_SERVER_CORS_ALLOWED_DOMAINS", "http://localhost:3000")

	fmt.Println("Token:", customAPIToken)
	fmt.Println("Server port:", serverPort)

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
			Timeout: 15 * time.Second,
		}),
		// Redirect trailing slash, cuz those are freaking cringe
		middleware.RemoveTrailingSlash(),
		customHeader,
	)

	// Routes START

	routes.RegisterBasicRoutes(e)
	routes.RegisterCharacterRoutes(e)

	// Routes END

	go func() {
		if err := e.Start(fmt.Sprintf(":%d", serverPort)); err != nil && err != http.ErrServerClosed {
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
