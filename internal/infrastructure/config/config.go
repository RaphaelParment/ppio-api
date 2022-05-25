package config

import "github.com/caarlos0/env/v6"

type Config struct {
	DB struct {
		User       string `env:"DB_USER" envDefault:"ppio"`
		Password   string `env:"DB_PASSWORD" envDefault:"dummy"`
		Host       string `env:"DB_HOST" envDefault:"0.0.0.0"`
		Name       string `env:"DB_NAME" envDefault:"ppio"`
		DisableTLS bool   `env:"DB_DISABLE_TLS" envDefault:"true"`
	}
}

func NewConfig() (*Config, error) {
	var config = &Config{}
	if err := env.Parse(config); err != nil {
		return nil, err
	}

	return config, nil
}
