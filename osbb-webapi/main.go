package main

import (
	"context"
	"time"

	"github.com/i3odja/osbb/shared/logger"
	"github.com/i3odja/osbb/webapi/client"
	"github.com/i3odja/osbb/webapi/config"
	"github.com/i3odja/osbb/webapi/server"
)

func main() {
	ctx := context.Background()
	logger := logger.NewLogger("osbb-webapi")

	logger.Infoln("Webapi service starting...")

	nc, err := config.NewConfig()
	if err != nil {
		logger.WithError(err).Infoln("could not get new config...")
	}

	address, err := nc.OSBBNotificationsConfig(ctx)
	if err != nil {
		logger.WithError(err).Infoln("could not get osbb-notifications config...")
	}

	c, err := client.NewNotifications(address.OSBBNotifications)
	if err != nil {
		logger.WithError(err).Fatalln("Could not create notifications client")
	}
	defer c.Close()

	// Contact the controller and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = server.AllNotifications(ctx, logger, c)
	if err != nil {
		logger.WithError(err).Fatalln("Could not sent all notifications")
	}
}
