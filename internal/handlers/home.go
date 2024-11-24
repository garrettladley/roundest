package handlers

import (
	"github.com/garrettladley/roundest/internal/templ"
	"github.com/garrettladley/roundest/internal/views/home"
	"github.com/gofiber/fiber/v2"
)

func (s *Service) Home(c *fiber.Ctx) error {
	pair, err := s.store.RandomPair(c.Context())
	if err != nil {
		return err
	}

	return templ.Render(c, home.Index(pair))
}
