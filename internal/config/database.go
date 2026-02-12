package config

type (
	DatabaseConfig struct {
		DatabaseDSN string `env:"DATABASE_DSN" env-required:"true" yaml:"dsn"`
	}

	DatabaseConfigInterface interface {
		GetDatabaseDSN() string
	}
)

func (d *DatabaseConfig) GetDatabaseDSN() string {
	return d.DatabaseDSN
}
