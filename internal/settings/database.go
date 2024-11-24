package settings

import "time"

type Database struct {
	DSN             string        `env:"URL" envDefault:"postgres://postgres:password@127.0.0.1:5432/roundest?sslmode=disable"`
	MaxOpenConns    uint          `env:"MAX_OPEN_CONNS" envDefault:"25"`
	MaxIdleConns    uint          `env:"MAX_IDLE_CONNS" envDefault:"25"`
	ConnMaxLifetime time.Duration `env:"CONN_MAX_LIFETIME" envDefault:"5m"`
}
