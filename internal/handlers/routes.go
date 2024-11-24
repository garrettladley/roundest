package handlers

import "github.com/gofiber/fiber/v2"

func (s *Service) Routes(r fiber.Router) {
	r.Get("/", s.Home)
	r.Post("/vote", s.Vote)
	r.Get("/results", s.Results)
}
