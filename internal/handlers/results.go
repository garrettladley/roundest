package handlers

import (
	"github.com/garrettladley/roundest/internal/templ"
	"github.com/garrettladley/roundest/internal/views/results"
	"github.com/gofiber/fiber/v2"
)

func (s *Service) Results(c *fiber.Ctx) error {
	xresults, err := s.store.GetAllResults(c.Context())
	if err != nil {
		return err
	}

	return templ.Render(c, results.Index(xresults))
}
