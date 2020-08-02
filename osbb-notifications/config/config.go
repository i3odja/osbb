package config

import (
	"context"
	"fmt"

	"github.com/i3odja/osbb/notifications/controller"
	"github.com/i3odja/osbb/notifications/storage"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PostgresHost string `envconfig:"POSTGRES_HOST" required:"true"`
	PostgresPort int    `envconfig:"POSTGRES_PORT" required:"true"`
	PostgresUser string `envconfig:"POSTGRES_USER" required:"true"`
	PostgresPass string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	PostgresDB   string `envconfig:"POSTGRES_DB" required:"true"`

	HTTPAddress      string `envconfig:"HTTP_ADDRESS" required:"true"`
	GRPCAddress      string `envconfig:"GRPC_ADDRESS" required:"true"`
	WebsocketAddress string `envconfig:"WEBSOCKET_ADDRESS" required:"true"`
}

// NewConfig() create new configuration for application
func NewConfig() (*Config, error) {
	var config Config
	// Read all environment variables and fill config structure with them
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, fmt.Errorf("envconfig error %w", err)
	}

	return &config, nil
}

// DBConfig get configuration for Postgres Database
func (c *Config) DBConfig(ctx context.Context) (*storage.DBConfig, error) {
	return &storage.DBConfig{
		Host:     c.PostgresHost,
		Port:     c.PostgresPort,
		User:     c.PostgresUser,
		Password: c.PostgresPass,
		DBName:   c.PostgresDB,
	}, nil
}

// Addresses get configuration for Postgres Database
func (c *Config) AddressConfig(ctx context.Context) (*controller.Address, error) {
	return &controller.Address{
		HTTP:      c.HTTPAddress,
		GRPC:      c.GRPCAddress,
		Websocket: c.WebsocketAddress,
	}, nil
}
