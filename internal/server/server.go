package server

import (
	"net/http"
	"strings"

	go_json "github.com/goccy/go-json"

	"github.com/garrettladley/roundest/internal/handlers"
	"github.com/garrettladley/roundest/internal/settings"
	"github.com/garrettladley/roundest/internal/storage/postgres"
	"github.com/garrettladley/roundest/internal/xerr"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func newApp(settings settings.Settings) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:  go_json.Marshal,
		JSONDecoder:  go_json.Unmarshal,
		ErrorHandler: xerr.ErrorHandler,
	})
	setupMiddleware(app)
	setupHealthCheck(app)
	setupFavicon(app)
	setupCaching(app)
	store, err := postgres.New(postgres.From(settings.Database))
	if err != nil {
		panic("internal/server: failed to create store: " + err.Error())
	}

	handlers.NewService(store).Routes(app)

	return app
}

func setupMiddleware(app *fiber.App) {
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}:${port} ${pid} ${locals:requestid} ${status} - ${latency} ${method} ${path}\n",
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join([]string{"https://garrettladley.com", "http://garrettladley.com", "http://127.0.0.1"}, ","),
		AllowMethods:     strings.Join([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions}, ","),
		AllowHeaders:     strings.Join([]string{"Accept", "Authorization", "Content-Type"}, ","),
		AllowCredentials: true,
		MaxAge:           300,
	}))
}

func setupHealthCheck(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})
}

func setupFavicon(app *fiber.App) {
	app.Get("/favicon.ico", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusNoContent)
	})
}
