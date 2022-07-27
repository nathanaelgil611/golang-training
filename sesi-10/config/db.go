package config

type ConfigDatabase struct {
	Port     string `env:"PORT"`
	Host     string `env:"HOST"`
	DBName   string `env:"DBNAME"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
}
