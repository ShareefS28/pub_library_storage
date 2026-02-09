package middlewares

import (
	"server/database"
	"server/models"
	"server/utils/utilresponse"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func BFFMiddlewares() fiber.Handler {
	return func(c fiber.Ctx) error {

		// session from cookie
		sessionUUIDstr := c.Cookies("session_uuid")
		if sessionUUIDstr == "" {
			return utilresponse.Error(
				c,
				"A00",
				fiber.StatusUnauthorized,
				"Session cookie not found",
			)
		}

		sessionUUID, err := uuid.Parse(sessionUUIDstr)
		if err != nil {
			return utilresponse.Error(
				c,
				"P01",
				fiber.StatusBadRequest,
				"Invalid UUID format",
			)
		}

		// Load session from DB
		var sess models.Session
		tx_sess := database.DB.Raw(`
			SELECT 
				CAST(uuid AS CHAR(36)) AS uuid,
				mt_account_id,
				CAST(created_by AS CHAR(36)) AS created_by,
				expired_at
			FROM dbo.mt_sessions
			WHERE uuid = ?
			  AND is_deleted = 0
			  AND revoked_at IS NULL
			  AND expired_at > SYSUTCDATETIME()
		`, sessionUUID).Scan(&sess)

		if tx_sess.Error != nil {
			return utilresponse.Error(
				c,
				"A99",
				fiber.StatusUnauthorized,
				tx_sess.Error.Error(),
			)
		}

		if tx_sess.RowsAffected == 0 {
			return utilresponse.Error(
				c,
				"A01",
				fiber.StatusUnauthorized,
				"Session not found or expired",
			)
		}

		// store user in context
		c.Locals("account_id", sess.AccountID)
		c.Locals("account_uuid", sess.CreatedBy)
		c.Locals("session_uuid", sess.UUID)
		c.Locals("expired_at", sess.ExpiredAt)

		// Continue request flow
		return c.Next()
	}
}
