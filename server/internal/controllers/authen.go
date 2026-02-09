package controllers

import (
	auth "server/internal/handlers/authen"
	reg "server/internal/handlers/register"

	"github.com/gofiber/fiber/v3"
)

func RegisterAuthen(api fiber.Router, protected fiber.Router) {
	api.Post("/auth/login", auth.Login)
	api.Post("/auth/register", reg.RegisterAccount)

	protected.Post("/auth/refreshSession", auth.RefreshToken)
	protected.Delete("/auth/logout", auth.Logout)
}
