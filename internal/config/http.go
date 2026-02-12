package config

import "time"

type (
	HTTPConfig struct {
		Port            string `env:"HTTP_PORT" env-default:"8000" yaml:"port"`
		ReadTimeout     string `env:"HTTP_READ_TIMEOUT" env-default:"120s" yaml:"read_timeout"`
		WriteTimeout    string `env:"HTTP_WRITE_TIMEOUT" env-default:"120s" yaml:"write_timeout"`
		ShutdownTimeout string `env:"HTTP_SHUTDOWN_TIMEOUT" env-default:"3s" yaml:"shutdown_timeout"`

		readTimeout     time.Duration `env:"-" yaml:"-"`
		writeTimeout    time.Duration `env:"-" yaml:"-"`
		shutdownTimeout time.Duration `env:"-" yaml:"-"`

		CORSOrigins []string `env:"HTTP_CORS_ORIGINS" env-separator:"," env-default:"*" yaml:"cors_origins"`
	}

	HTTPConfigInterface interface {
		GetCORSOrigins() []string
		GetPort() string
		GetReadTimeout() time.Duration
		GetWriteTimeout() time.Duration
		GetShutdownTimeout() time.Duration
	}
)

func (h *HTTPConfig) GetPort() string {
	return h.Port
}

func (h *HTTPConfig) GetReadTimeout() time.Duration {

	if h.readTimeout == 0 {
		var err error
		h.readTimeout, err = time.ParseDuration(h.ReadTimeout)

		if err != nil {
			h.readTimeout = 120 * time.Second
		}
	}

	return h.readTimeout
}

func (h *HTTPConfig) GetWriteTimeout() time.Duration {
	if h.writeTimeout == 0 {
		var err error
		h.writeTimeout, err = time.ParseDuration(h.WriteTimeout)

		if err != nil {
			h.writeTimeout = 120 * time.Second
		}
	}

	return h.writeTimeout
}

func (h *HTTPConfig) GetShutdownTimeout() time.Duration {
	if h.shutdownTimeout == 0 {
		var err error
		h.shutdownTimeout, err = time.ParseDuration(h.ShutdownTimeout)

		if err != nil {
			h.shutdownTimeout = 120 * time.Second
		}
	}

	return h.shutdownTimeout
}

func (h *HTTPConfig) GetCORSOrigins() []string {
	return h.CORSOrigins
}
