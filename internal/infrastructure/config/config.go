package config

import "github.com/caarlos0/env/v6"

type Config struct {
	DB struct {
		User       string `env:"DB_USER" envDefault:"ppio"`
		Password   string `env:"DB_PASSWORD" envDefault:"dummy"`
		Host       string `env:"DB_HOST" envDefault:"localhost"`
		Port       string `env:"DB_PORT" envDefault:":5432"`
		Name       string `env:"DB_NAME" envDefault:"ppio"`
		DisableTLS bool   `env:"DB_DISABLE_TLS" envDefault:"true"`
	}

	Http struct {
		Port string `env:"HTTP_PORT" envDefault:":9001"`
	}
}

func NewConfig() (*Config, error) {
	var config = &Config{}
	if err := env.Parse(config); err != nil {
		return nil, err
	}

	return config, nil
}
