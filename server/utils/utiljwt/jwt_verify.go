package utiljwt

import (
	"server/config"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(accessToken string) (*jwt.Token, error) {
	return jwt.Parse(
		accessToken,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return publicKey, nil
		},
		jwt.WithIssuer(config.AppConfig.JWTIssuer),
		jwt.WithValidMethods([]string{"ES256"}),
	)
}
