package config

import "time"

type ConfigMiddleware struct {
	Username string `env:"USERNAME_MW"`
	Password string `env:"PASSWORD_MW"`
}

type ConfigJWT struct {
	ApplicationName         string        `env:"APPLICATION_NAME"`
	LoginExpirationDuration time.Duration `env:"LOGIN_EXPIRATION_DURATION"`
	SigningMethod           string        `env:"JWT_SIGNING_METHOD"`
	SignatureKey            []byte        `env:"JWT_SIGNATURE_KEY"`
}
