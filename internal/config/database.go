package config

import "fmt"

type DatabaseConfig struct {
	Database string `env:"DATABASE_DB"`
	Username string `env:"DATABASE_USER"`
	Password string `env:"DATABASE_PASS"`
	Port     int    `env:"DB_PORT" envDefault:"5432"`
	Host     string `env:"DATABASE_HOST" envDefault:"localhost"`
}

func (cfg *DatabaseConfig) DatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
}
