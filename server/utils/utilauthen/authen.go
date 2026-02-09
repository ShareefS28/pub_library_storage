package utilauthen

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func HashValue(value string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(value),
		bcrypt.DefaultCost,
	)

	return string(bytes), err
}

func CheckHash(hash_value string, value string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash_value), []byte(value))
}

func RandomString(size int) (string, error) {
	b := make([]byte, size)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}
