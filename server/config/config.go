package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var AppConfig *Config

type Config struct {
	AppEnv  string
	AppName string
	AppPort string

	DBDSN string

	JWTPrivateKey []byte
	JWTPublicKey  []byte

	JWTAccessTTLMin  int
	JWTRefreshTTLDAY int
	JWTIssuer        string

	IsProd bool
}

func Load() *Config {
	// Load .env only in local/dev
	_ = godotenv.Load()

	privateKeyPath := getEnv("JWT_PRIVATE_KEY_PATH", "")
	publicKeyPath := getEnv("JWT_PUBLIC_KEY_PATH", "")

	privateKey, err := os.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("failed to read private key: %v", err)
	}

	publicKey, err := os.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatalf("failed to read public key: %v", err)
	}

	cfg := &Config{
		AppEnv:  getEnv("APP_ENV", "DEV"),
		AppName: getEnv("APP_NAME", "fiber-auth"),
		AppPort: getEnv("APP_PORT", "3000"),

		DBDSN: getEnv("DB_DSN", ""),

		JWTPrivateKey: privateKey,
		JWTPublicKey:  publicKey,

		JWTAccessTTLMin:  getEnvAsInt("JWT_ACCESS_TTL_MIN", 15),
		JWTRefreshTTLDAY: getEnvAsInt("JWT_REFRESH_TTL_DAY", 7),
		JWTIssuer:        getEnv("JWT_ISSUER", "library_storage"),

		IsProd: getEnvAsBool("IsProd", false),
	}

	validate(cfg)

	log.Println("Config Load Success")

	return cfg
}

func validate(cfg *Config) {
	if cfg.DBDSN == "" {
		log.Fatal("DB_DSN is required")
	}
	if len(cfg.JWTPrivateKey) == 0 || len(cfg.JWTPublicKey) == 0 {
		log.Fatal("JWT key paths are required")
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}

	return fallback
}
