package conf

import (
	"github.com/caarlos0/env/v6"
)

// AppConfig presents app conf
type AppConfig struct {
	AppEnv   string `env:"APP_ENV" envDefault:"dev"`
	Port     string `env:"PORT" envDefault:"8000"`
	DBHost   string `env:"DB_HOST" envDefault:"localhost"`
	DBPort   string `env:"DB_PORT" envDefault:"5432"`
	DBUser   string `env:"DB_USER" envDefault:"postgres"`
	DBPass   string `env:"DB_PASS" envDefault:"postgres"`
	DBName   string `env:"DB_NAME" envDefault:"service_user"`
	EnableDB string `env:"ENABLE_DB" envDefault:"true"`

	API_SECRET string `env:"API_SECRET" envDefault:"as@NCIQ3A0!"`
}

var config AppConfig

func SetEnv() {
	_ = env.Parse(&config)
}

func LoadEnv() AppConfig {
	return config
}
