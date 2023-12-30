package config

import (
	"time"
)

type Config struct {
	Env        string `env:"env" env-default:"local"`
	HTTPServer `env:"http_server"`
	Database   `env:"database"`
}

type HTTPServer struct {
	Address     string        `env:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `env:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `env:"idle_timeout" env-default:"60s"`
	User        string        `env:"user" env-required:"true"`
	Password    string        `env:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

type Database struct {
	DB        string `env:"db" env-default:"postgres"`
	DBUser    string `env:"db_user" env-default:"postgres"`
	DBPass    string `env:"db_pass" env-default:"secret"`
	DBPort    int64  `env:"db_port" env-default:"5432"`
	DBHost    string `env:"db_host" env-default:"localhost"`
	DBSslmode string `env:"db_sslmode" env-default:"sslmode"`
}
