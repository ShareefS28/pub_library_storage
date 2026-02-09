package controllers

import (
	prof "server/internal/handlers/profile"

	"github.com/gofiber/fiber/v3"
)

func RegisterProfile(protected fiber.Router) {
	protected.Get("/me", prof.Me)
}
