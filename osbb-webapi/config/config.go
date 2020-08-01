package config

import (
	"context"
	"fmt"

	"github.com/i3odja/osbb/webapi/client"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	OSBBNotificationsAddress string `envconfig:"OSBB_NOTIFICATIONS_ADDRESS" required:"true"`
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

// OSBBNotificationsConfig() get configuration for webapi client
func (c *Config) OSBBNotificationsConfig(ctx context.Context) (*client.Address, error) {
	return &client.Address{OSBBNotifications: c.OSBBNotificationsAddress}, nil
}
