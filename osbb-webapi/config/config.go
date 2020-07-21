package config

import "time"

type Config struct {
	HTTPIP       string        `envconfig:"HTTP_IP" default:"0.0.0.0"`
	HTTPPort     int           `envconfig:"HTTP_PORT" default:"9999"`
	ReadTimeout  time.Duration `envconfig:"READ_TIMEOUT" default:"60s"`
	WriteTimeout time.Duration `envconfig:"WRITE_TIMEOUT" default:"60s"`
}
