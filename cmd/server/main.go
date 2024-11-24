package main

import (
	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/garrettladley/roundest/internal/server"
	"github.com/garrettladley/roundest/internal/settings"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	settings, err := settings.Load()
	if err != nil {
		log.Fatalf("Failed to load settings: %v", err)
	}

	app := server.Init(settings)

	static(app)

	go func() {
		if err := app.Listen(":" + settings.App.Port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	slog.Info("Shutting down server")
	if err := app.Shutdown(); err != nil {
		slog.Error("failed to shutdown server", "error", err)
	}

	slog.Info("Server shutdown")
}

//go:embed public
var PublicFS embed.FS

//go:embed deps
var DepsFS embed.FS

func static(app *fiber.App) {
	app.Use("/public", filesystem.New(filesystem.Config{
		Root:       http.FS(PublicFS),
		PathPrefix: "public",
		Browse:     true,
	}))
	app.Use("/deps", filesystem.New(filesystem.Config{
		Root:       http.FS(DepsFS),
		PathPrefix: "deps",
		Browse:     true,
	}))
}
