package config

type ConfigMiddleware struct {
	Username string `env:"USERNAME_MW"`
	Password string `env:"PASSWORD_MW"`
}
