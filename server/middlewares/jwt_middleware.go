package middlewares

import (
	"server/config"
	"server/database"
	"server/utils/utiljwt"
	"server/utils/utilresponse"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddlewares() fiber.Handler {
	return func(c fiber.Ctx) error {

		// access_token from cookie
		accessToken := c.Cookies("access_token")
		if accessToken == "" {
			return utilresponse.Error(
				c,
				"A01",
				fiber.StatusUnauthorized,
				"Unauthorized",
			)
		}

		// verify JWT
		token, err := utiljwt.VerifyToken(accessToken)
		if err != nil || !token.Valid {
			return utilresponse.Error(
				c,
				"A02",
				fiber.StatusUnauthorized,
				"Invalid or expired token",
			)
		}

		// Extract claims, check is type match with jwt.MapClaims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return utilresponse.Error(
				c,
				"A03",
				fiber.StatusUnauthorized,
				"Invalid token claims",
			)
		}

		// check sessions Id
		sid, ok := claims["sid"].(string)
		if !ok {
			return utilresponse.Error(
				c,
				"A04",
				fiber.StatusUnauthorized,
				"Invalid session",
			)
		}

		// check revoke session
		var exists bool
		tx_sess := database.DB.Raw(`
			SELECT CASE 
			WHEN EXISTS (
					SELECT 1
					FROM dbo.mt_sessions
					WHERE uuid = ?
					AND (revoked_at IS NOT NULL OR is_deleted = 1 OR expired_at < SYSUTCDATETIME()) 
				) THEN 1 ELSE 0
			END
		`, sid).Scan(&exists)

		if tx_sess.Error != nil || exists {
			return utilresponse.Error(
				c,
				"A05",
				401,
				"Session Revoked Or Deleted",
			)
		}

		// get account uuid
		sub, ok := claims["sub"].(string)
		if !ok {
			return utilresponse.Error(
				c,
				"A06",
				fiber.StatusUnauthorized,
				"Invalid subject",
			)
		}

		// get issuer
		iss, ok := claims["iss"].(string)
		if !ok || iss != config.AppConfig.JWTIssuer {
			return utilresponse.Error(
				c,
				"A05",
				fiber.StatusUnauthorized,
				"Invalid issuer",
			)
		}

		// check expire jwt
		exp, ok := claims["exp"].(float64)
		if !ok || int64(exp) < time.Now().UTC().Unix() {
			return utilresponse.Error(
				c,
				"A07",
				fiber.StatusUnauthorized,
				"Token Expired Or Invalid exp claim",
			)
		}

		// store user in context
		c.Locals("access_token", accessToken)
		c.Locals("refresh_plain", accessToken)
		c.Locals("account_uuid", sub)
		c.Locals("session_uuid", sid)
		c.Locals("issuer", iss)

		// Continue request flow
		return c.Next()
	}
}
