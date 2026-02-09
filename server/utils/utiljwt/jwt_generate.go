package utiljwt

import (
	"server/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateAccessToken(accountUUID uuid.UUID, sessionUUID uuid.UUID, ttlMinutes int) (string, error) {
	claims := jwt.MapClaims{
		"alg": "ES256",
		"typ": "JWT",
		"sid": sessionUUID,
		"sub": accountUUID,
		"exp": time.Now().UTC().Add(time.Minute * time.Duration(ttlMinutes)).Unix(),
		"iat": time.Now().UTC().Unix(),
		"iss": config.AppConfig.JWTIssuer,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	return token.SignedString(privateKey)
}
