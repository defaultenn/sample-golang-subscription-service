package config

import (
	"errors"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	AppConfig struct {
		DatabaseConfig DatabaseConfig `yaml:"database"`
		HTTPConfig     HTTPConfig     `yaml:"http"`
		LogConfig      LogConfig      `yaml:"logging"`
	}

	AppConfigInterface interface {
		GetDatabaseConfig() DatabaseConfigInterface
		GetHTTPConfig() HTTPConfigInterface
		GetLogConfig() LogConfigInterface
	}
)

func NewConfig() AppConfigInterface {
	cfg := &AppConfig{}

	if err := cleanenv.ReadConfig("./config.yaml", cfg); err != nil && !errors.Is(err, os.ErrNotExist) {
		panic(err)
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		panic(err)
	}

	return cfg
}

func (d *AppConfig) GetDatabaseConfig() DatabaseConfigInterface {
	return &d.DatabaseConfig
}

func (d *AppConfig) GetHTTPConfig() HTTPConfigInterface {
	return &d.HTTPConfig
}

func (d *AppConfig) GetLogConfig() LogConfigInterface {
	return &d.LogConfig
}
