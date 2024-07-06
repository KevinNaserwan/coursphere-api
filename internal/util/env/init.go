package env

import (
	"github.com/caarlos0/env/v8"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Environment   string `env:"APP_ENV,unset"                       envDefault:"debug"`
	Port          int16  `env:"APP_PORT,unset"                      envDefault:"5000"`
	JwtSecret     string `env:"JWT_SECRET,unset"`
	DbReleaseURL  string `env:"POSTGRES_CONNECTION_URL,unset"`
	DbTestURL     string `env:"POSTGRES_TEST_CONNECTION_URL,unset"`
	DbDebugURL    string `env:"POSTGRES_DEBUG_CONNECTION_URL,unset"`
	AdminUsername string `env:"ADMIN_USERNAME"`
	AdminPassword string `env:"ADMIN_PASSWORD"`
	EmailUser     string `env:"EMAIL_USER,unset"`
	EmailFrom     string `env:"EMAIL_FROM"`
	EmailPassword string `env:"EMAIL_PASSWORD,unset"`
	EmailHost     string `env:"EMAIL_HOST,unset"`
	EmailPort     int16  `env:"EMAIL_PORT,unset"`
	FrontendURL   string `env:"FRONTEND_URL"`
}

func LoadConfig() *Config {
	cfg := new(Config)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
