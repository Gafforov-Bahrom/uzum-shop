package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

/*
 */
type Config struct {
	DatabasePassword string `envconfig:"DB_PASSWORD"`
	DatabaseUsername string `envconfig:"DB_USER"`
	DatabaseName     string `envconfig:"DB_NAME"`
	DatabaseHost     string `envconfig:"DB_HOST"`
	DatabasePort     int64  `envconfig:"DB_PORT"`
	Port             string `envconfig:"PORT"`
	HTTPPort         string `envconfig:"HTTP_PORT"`
}

func NewConfig(envSrc string) (*Config, error) {
	cfg := &Config{}
	err := godotenv.Load(envSrc)
	envconfig.Process("", cfg)
	return cfg, err
}
