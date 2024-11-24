//go:build !dev

package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/utils"
)

func setupCaching(app *fiber.App) {
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			_, ok := staticPaths[c.OriginalURL()]
			return !ok
		},
		KeyGenerator: func(c *fiber.Ctx) string { return utils.CopyString(c.OriginalURL()) },
		Expiration:   time.Hour * 24 * 365, // 1 year
		CacheControl: true,
	}))
}
