package controllers

import (
	"server/utils/utilresponse"

	"github.com/gofiber/fiber/v3"
)

// @Summary      Health check
// @Description Check if server is running
// @Tags         Health
// @Accept       json
// @Produce      json
// @Router       /health [get]
func RegisterHealth(api fiber.Router) {
	api.Get("/", func(c fiber.Ctx) error {
		return utilresponse.Success(
			c,
			fiber.StatusOK,
			fiber.Map{
				"message": "Server Already Running",
			},
		)
	})

	api.Get("/health", func(c fiber.Ctx) error {
		return utilresponse.Success(
			c,
			fiber.StatusOK,
			fiber.Map{
				"message": "health check ok",
			},
		)
	})
}

// @Summary      Health check
// @Description Check if server is running
// @Tags         Health
// @Accept       json
// @Produce      json
// @Router       /secure/health [get]
func ProtectedRegisterHealth(protected fiber.Router) {
	protected.Get("/health", func(c fiber.Ctx) error {
		return utilresponse.Success(
			c,
			fiber.StatusOK,
			fiber.Map{
				"message": "health check ok",
			},
		)
	})
}
