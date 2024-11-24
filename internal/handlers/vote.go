package handlers

import (
	"github.com/garrettladley/roundest/internal/templ"
	"github.com/garrettladley/roundest/internal/views/home"
	"github.com/garrettladley/roundest/internal/xerr"
	"github.com/gofiber/fiber/v2"
)

type voteRequest struct {
	UpID   int `query:"up"`
	DownID int `query:"down"`
}

func (s *Service) Vote(c *fiber.Ctx) error {
	var req voteRequest
	if err := c.QueryParser(&req); err != nil {
		return xerr.InvalidJSON(err)
	}

	if err := s.store.Vote(c.Context(), req.UpID, req.DownID); err != nil {
		return err
	}

	newPair, err := s.store.RandomPair(c.Context())
	if err != nil {
		return err
	}

	return templ.Render(c, home.Ballot(newPair))
}
