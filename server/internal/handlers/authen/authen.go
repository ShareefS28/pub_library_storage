package authen

import (
	"errors"
	"server/config"
	"server/database"
	"server/dtos/dtologin"
	"server/models"
	"server/utils/utilauthen"
	"server/utils/utilresponse"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// @Tags         Authentication
// @Summary      Generate Cookie
// @Description  Generate Cookie session
// @Accept       json
// @Produce      json
// @Param		 request body dtologin.DTOLoginAccountReq true "Login credentials"
// @Router       /auth/login [post]
func Login(c fiber.Ctx) error {

	conn := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			conn.Rollback()
		}
	}()

	// check body of request
	var req dtologin.DTOLoginAccountReq
	if err := c.Bind().Body(&req); err != nil {
		return utilresponse.Error(
			c,
			"L00",
			fiber.StatusBadRequest,
			"Invalid request body",
		)
	}

	// check account
	var acc models.Account
	tx_acc := conn.Raw(`
		SELECT id, CAST(uuid AS CHAR(36)) AS uuid, username, password_hash
		FROM dbo.mt_accounts
		WHERE username = ?
			AND is_deleted = 0
	`, req.Username).First(&acc)

	if tx_acc.Error != nil {

		if errors.Is(tx_acc.Error, gorm.ErrRecordNotFound) {
			return utilresponse.Error(
				c,
				"L02",
				fiber.StatusUnauthorized,
				"Username Not Found",
			)
		}

		return utilresponse.Error(
			c,
			"D00",
			fiber.StatusInternalServerError,
			"Database error",
		)
	}

	if tx_acc.RowsAffected == 0 {
		return utilresponse.Error(
			c,
			"L01",
			fiber.StatusUnauthorized,
			"Username Or Password Is Wrong",
		)
	}

	if err := utilauthen.CheckHash(acc.PasswordHash, req.Password); err != nil {
		return utilresponse.Error(
			c,
			"L02",
			fiber.StatusUnauthorized,
			"Username Or Password Is Wrong",
		)
	}

	// create session
	if err := createdSessionToken(c, acc, conn); err != nil {
		if e, ok := err.(*utilresponse.AppError); ok {
			return utilresponse.EnrollErrorRollback(c, e, conn)
		}

		return utilresponse.Error(c, "GEN99", 500, err.Error())
	}

	if err := conn.Commit().Error; err != nil {
		conn.Rollback()
		return utilresponse.Error(
			c,
			"T99",
			fiber.StatusInternalServerError,
			"Transaction commit failed",
		)
	}

	res := dtologin.DTOLoginSuccessRes{}
	res.User.UUID = acc.UUID
	res.User.Username = acc.Username

	return utilresponse.Success(
		c,
		fiber.StatusOK,
		res,
	)
}

// @Tags         Authentication
// @Summary      Destroy Cookie
// @Description  Destroy Cookie session
// @Accept       json
// @Produce      json
// @Router       /secure/auth/logout [delete]
func Logout(c fiber.Ctx) error {

	conn := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			conn.Rollback()
		}
	}()

	// check account
	accountUUID, ok := c.Locals("account_uuid").(uuid.UUID)
	if !ok {
		return utilresponse.Error(
			c,
			"P01",
			fiber.StatusBadRequest,
			"Invalid UUID format",
		)
	}

	// Delete Session
	if err := deletedSessionToken(c, accountUUID, conn); err != nil {
		if e, ok := err.(*utilresponse.AppError); ok {
			return utilresponse.EnrollErrorRollback(c, e, conn)
		}

		return utilresponse.Error(c, "GEN99", 500, err.Error())
	}

	if err := conn.Commit().Error; err != nil {
		conn.Rollback()
		return utilresponse.Error(
			c,
			"T99",
			fiber.StatusInternalServerError,
			"Transaction commit failed",
		)
	}

	return utilresponse.Success(
		c,
		fiber.StatusOK,
		fiber.Map{
			"message": "logout success",
		},
	)
}

// @Tags         Authentication
// @Summary      Refresh Cookie
// @Description  Refresh Cookie session
// @Accept       json
// @Produce      json
// @Router       /secure/auth/refreshSession [post]
func RefreshToken(c fiber.Ctx) error {

	conn := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			conn.Rollback()
		}
	}()

	// // get session of refresh_token cookies
	// refreshPlain := c.Cookies("refresh_token")
	// if refreshPlain == "" {
	// 	return utilresponse.Error(
	// 		c,
	// 		"R00",
	// 		fiber.StatusUnauthorized,
	// 		"Refresh Token Not Found",
	// 	)
	// }

	// Find session
	var session models.Session
	sessionUUID, ok := c.Locals("session_uuid").(uuid.UUID)
	if !ok {
		return utilresponse.Error(
			c,
			"A04",
			fiber.StatusUnauthorized,
			"Invalid Session",
		)
	}

	tx_sess := conn.Raw(`
		SELECT CAST(uuid AS CHAR(36)) AS uuid
		FROM dbo.mt_sessions
		WHERE uuid = ?
	`, sessionUUID).First(&session)

	if tx_sess.Error != nil {
		return utilresponse.Error(
			c,
			"D00",
			fiber.StatusInternalServerError,
			"Database error",
		)
	}

	if tx_sess.RowsAffected == 0 {
		return utilresponse.Error(
			c,
			"A04",
			fiber.StatusUnauthorized,
			"Invalid Session",
		)
	}

	// // Check Hash Refresh Token
	// if err := utilauthen.CheckHash(session.RefreshHashToken, refreshPlain); err != nil {
	// 	return utilresponse.Error(
	// 		c,
	// 		"A08",
	// 		fiber.StatusUnauthorized,
	// 		"Invalid refresh token",
	// 	)
	// }

	// check account
	accountUUID, ok := c.Locals("account_uuid").(uuid.UUID)
	if !ok {
		return utilresponse.Error(
			c,
			"P01",
			fiber.StatusBadRequest,
			"Invalid UUID format",
		)
	}

	var acc models.Account
	tx_acc := conn.Raw(`
		SELECT 
			id,
			username,
			CAST(uuid AS CHAR(36)) AS uuid
		FROM dbo.mt_accounts
		WHERE uuid = ?
			AND is_deleted = 0
	`, accountUUID).First(&acc)

	if tx_acc.Error != nil {
		return utilresponse.Error(
			c,
			"D00",
			fiber.StatusInternalServerError,
			"Database error",
		)
	}

	if tx_acc.RowsAffected == 0 {
		return utilresponse.Error(
			c,
			"A09",
			fiber.StatusUnauthorized,
			"Account Not Found",
		)
	}

	// Delete Session
	if err := deletedSessionToken(c, accountUUID, conn); err != nil {
		if e, ok := err.(*utilresponse.AppError); ok {
			return utilresponse.EnrollErrorRollback(c, e, conn)
		}

		return utilresponse.Error(c, "GEN99", 500, err.Error())
	}

	// create session
	if err := createdSessionToken(c, acc, conn); err != nil {
		if e, ok := err.(*utilresponse.AppError); ok {
			return utilresponse.EnrollErrorRollback(c, e, conn)
		}

		return utilresponse.Error(c, "GEN99", 500, err.Error())
	}

	if err := conn.Commit().Error; err != nil {
		conn.Rollback()
		return utilresponse.Error(
			c,
			"T99",
			fiber.StatusInternalServerError,
			"Transaction commit failed",
		)
	}

	expiredAt, ok := c.Locals("expired_at").(time.Time)
	if !ok {
		return utilresponse.Error(
			c,
			"P01",
			fiber.StatusBadRequest,
			"Invalid Time format",
		)
	}

	res := dtologin.DTOLoginSuccessRes{}
	res.User.UUID = acc.UUID
	res.User.Username = acc.Username
	res.User.ExpiredAt = expiredAt

	return utilresponse.Success(
		c,
		fiber.StatusOK,
		res,
	)
}

//#region Private Function

func createdSessionToken(c fiber.Ctx, acc models.Account, conn *gorm.DB) error {
	// // Generate Refresh Token
	// refreshPlain, err := utilauthen.RandomString(64)
	// if err != nil {
	// 	return utilresponse.Error(
	// 		c,
	// 		"Z98",
	// 		fiber.StatusInternalServerError,
	// 		"Error To Generate Refresh Token",
	// 	)
	// }
	// hash := sha256.Sum256([]byte(refreshPlain))
	// refreshHash := hex.EncodeToString(hash[:])

	// Create Session
	ip := c.IP() // Client IP address
	ua := c.Get("User-Agent")
	exp := time.Now().UTC().Add(time.Duration(config.AppConfig.JWTAccessTTLMin) * time.Minute)
	sid := uuid.New() // session_uuid

	// // Generate JWT Token
	// accessToken, err := utiljwt.GenerateAccessToken(acc.UUID, sid, config.AppConfig.JWTAccessTTLMin)
	// if err != nil {
	// 	return &utilresponse.AppError{
	// 		ErrorCode:  "Z99",
	// 		StatusCode: fiber.StatusInternalServerError,
	// 		Message:    "Error To Generate Access Token",
	// 	}
	// }

	session := models.Session{
		BaseModel: models.BaseModel{
			CreatedBy: acc.UUID,
		},
		UUID:      sid,
		IPAddress: &ip,
		AccountID: acc.ID,
		ExpiredAt: exp,
		UserAgent: &ua,
	}

	if err := conn.Create(&session).Error; err != nil {
		return &utilresponse.AppError{
			ErrorCode:  "D01",
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Database Error",
		}
	}

	if err := createdCookies(c, sid); err != nil {
		return &utilresponse.AppError{
			ErrorCode:  "Z99",
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Error To Created Cookie",
		}
	}

	return nil
}

func deletedSessionToken(c fiber.Ctx, accountUUID uuid.UUID, conn *gorm.DB) error {
	sessionUUIDStr := c.Cookies("session_uuid")
	if sessionUUIDStr == "" {
		return &utilresponse.AppError{
			ErrorCode:  "D05",
			StatusCode: fiber.StatusBadRequest,
			Message:    "Session Not Found",
		}
	}
	sessionUUID, err := uuid.Parse(sessionUUIDStr)
	if err != nil {
		return utilresponse.Error(
			c,
			"P01",
			fiber.StatusBadRequest,
			"Invalid UUID format",
		)
	}

	updated_session := conn.Exec(`
		UPDATE dbo.mt_sessions
		SET 
			revoked_at = SYSUTCDATETIME(),
			is_deleted = 1,
			deleted_at = SYSUTCDATETIME(),
			deleted_by = ?
		WHERE uuid = ?
	`, accountUUID, sessionUUID)

	if updated_session.Error != nil {
		return &utilresponse.AppError{
			ErrorCode:  "D03",
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Database Error",
		}
	}

	if updated_session.RowsAffected == 0 {
		return &utilresponse.AppError{
			ErrorCode:  "D04",
			StatusCode: fiber.StatusNotFound,
			Message:    "Session not found",
		}
	}

	if err := deletedCookies(c); err != nil {
		return &utilresponse.AppError{
			ErrorCode:  "D00",
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Error To Delete Cookie Session",
		}
	}

	return nil
}

func createdCookies(c fiber.Ctx, sessionUUID uuid.UUID) error {
	sameSite := func() string {
		if config.AppConfig.IsProd {
			return fiber.CookieSameSiteNoneMode
		}

		return fiber.CookieSameSiteLaxMode
	}

	// Set Cookies
	c.Cookie(&fiber.Cookie{
		Name:     "session_uuid",
		Value:    sessionUUID.String(),
		Path:     "/",
		HTTPOnly: true,
		Secure:   config.AppConfig.IsProd,
		SameSite: sameSite(),
		Expires:  time.Now().UTC().Add(time.Duration(config.AppConfig.JWTAccessTTLMin) * time.Minute),
	})

	// c.Cookie(&fiber.Cookie{
	// 	Name:     "refresh_token",
	// 	Value:    refreshPlain,
	// 	Path:     "/",
	// 	HTTPOnly: true,
	// 	Secure:   config.AppConfig.IsProd,
	// 	SameSite: sameSite(),
	// 	Expires:  time.Now().UTC().Add(time.Duration(config.AppConfig.JWTRefreshTTLDAY) * 24 * time.Hour),
	// })

	return nil
}

func deletedCookies(c fiber.Ctx) error {

	// Clear Cookies
	c.Cookie(&fiber.Cookie{
		Name:   "session_uuid",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	// c.Cookie(&fiber.Cookie{
	// 	Name:   "refresh_token",
	// 	Value:  "",
	// 	Path:   "/",
	// 	MaxAge: -1,
	// })

	return nil
}
