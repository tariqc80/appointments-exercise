package config

import (
	"log"

	env "github.com/Netflix/go-env"
)

// Config stuct to store the application configuration
type Config struct {
	DatabaseHost     string `env:"DATABASE_HOST"`
	DatabasePort     string `env:"DATABASE_PORT"`
	DatabaseName     string `env:"DATABASE_NAME"`
	DatabaseUser     string `env:"DATABASE_USER"`
	DatabasePassword string `env:"DATABASE_PASSWORD"`
}

// ParseFromEnv reads env var into config struct
func ParseFromEnv() *Config {
	cfg := Config{}

	if _, err := env.UnmarshalFromEnviron(&cfg); err != nil {
		log.Print(err)
	}

	return &cfg
}
