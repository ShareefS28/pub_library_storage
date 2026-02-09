package controllers

import (
	book "server/internal/handlers/book"

	"github.com/gofiber/fiber/v3"
)

func RegisterBook(protected fiber.Router) {
	protected.Post("/book/create", book.CreateBook)
}
