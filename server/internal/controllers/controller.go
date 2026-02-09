package controllers

import (
	"server/middlewares"

	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	/**
	* Base API
	**/
	api := app.Group("/api")

	/**
	* Protected API After Has sessions
	**/
	protected := api.Group("/secure", middlewares.BFFMiddlewares())

	// Health
	RegisterHealth(api)
	ProtectedRegisterHealth(protected)

	// Authentication
	RegisterAuthen(api, protected)

	// Profile
	RegisterProfile(protected)

	// Book
	RegisterBook(protected)

}
