package utilsession

import (
	"server/utils/utilresponse"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SessionInfo struct {
	AccountID   int       `json:"account_id"`
	AccountUUID uuid.UUID `json:"account_uuid"`
	SessionUUID uuid.UUID `json:"session_uuid"`
	ExpiredAt   time.Time `json:"expired_at"`
}

func GetSessionInfo(c fiber.Ctx) (SessionInfo, error) {
	accountID, ok := c.Locals("account_id").(int)
	if !ok {
		return SessionInfo{}, &utilresponse.AppError{
			ErrorCode:  "P01",
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid UUID format",
		}
	}

	accountUUID, ok := c.Locals("account_uuid").(uuid.UUID)
	if !ok {
		return SessionInfo{}, &utilresponse.AppError{
			ErrorCode:  "P01",
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid UUID format",
		}
	}

	sessionUUID, ok := c.Locals("session_uuid").(uuid.UUID)
	if !ok {
		return SessionInfo{}, &utilresponse.AppError{
			ErrorCode:  "P01",
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid UUID format",
		}
	}

	expiredAt, ok := c.Locals("expired_at").(time.Time)
	if !ok {
		return SessionInfo{}, &utilresponse.AppError{
			ErrorCode:  "P01",
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid UUID format",
		}
	}

	return SessionInfo{
		AccountID:   accountID,
		AccountUUID: accountUUID,
		SessionUUID: sessionUUID,
		ExpiredAt:   expiredAt,
	}, nil
}

func FindSession(sessionUUID uuid.UUID, conn *gorm.DB) (bool, error) {
	var exists bool
	tx_sess := conn.Raw(`
		SELECT CASE 
			WHEN NOT EXISTS (
				SELECT 1
				FROM dbo.mt_sessions
				WHERE uuid = ?
					AND revoked_at IS NULL 
					AND is_deleted = 0 
					AND expired_at > SYSUTCDATETIME()
			) THEN 0 ELSE 1
		END
	`, sessionUUID).Scan(&exists)

	if tx_sess.Error != nil {
		return exists, &utilresponse.AppError{
			ErrorCode:  "D03",
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Database Error",
		}
	}

	if tx_sess.RowsAffected == 0 {
		return exists, &utilresponse.AppError{
			ErrorCode:  "D04",
			StatusCode: fiber.StatusNotFound,
			Message:    "Session not found",
		}
	}

	return exists, nil
}
