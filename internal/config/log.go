package config

import (
	"strings"

	"github.com/rs/zerolog"
)

type (
	LogConfig struct {
		Level string `env:"LOG_LEVEL" env-default:"error" yaml:"level"`
	}

	LogConfigInterface interface {
		GetLevel() zerolog.Level
	}
)

func (l *LogConfig) GetLevel() zerolog.Level {

	var lvl zerolog.Level

	switch strings.ToLower(l.Level) {
	case "error":
		lvl = zerolog.ErrorLevel
	case "warn":
		lvl = zerolog.WarnLevel
	case "info":
		lvl = zerolog.InfoLevel
	case "debug":
		lvl = zerolog.DebugLevel
	default:
		lvl = zerolog.InfoLevel
	}

	return lvl
}
