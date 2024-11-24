package server

import (
	"github.com/garrettladley/roundest/internal/settings"
	"github.com/gofiber/fiber/v2"
)

func Init(settings settings.Settings) *fiber.App {
	return newApp(settings)
}
