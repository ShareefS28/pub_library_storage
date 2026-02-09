package utilresponse

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type AppError struct {
	ErrorCode  string
	StatusCode int
	Message    string
}

func (e *AppError) Error() string {
	return e.Message
}

// RollBack and Error
func EnrollErrorRollback(c fiber.Ctx, err *AppError, conn *gorm.DB) error {
	conn.Rollback()
	return Error(c, err.ErrorCode, err.StatusCode, err.Message)
}

func Error(c fiber.Ctx, error_code string, status_code int, message string) error {
	return c.Status(status_code).JSON(fiber.Map{
		"status_code": status_code,
		"error_code":  error_code,
		"success":     false,
		"message":     message,
	})
}

func Success(c fiber.Ctx, status_code int, data interface{}) error {
	return c.Status(status_code).JSON(fiber.Map{
		"status_code": status_code,
		"success":     true,
		"data":        data,
	})
}
