package main

import (
	"context"
	"os"
	"time"

	"github.com/i3odja/osbb/shared/logger"
	"github.com/i3odja/osbb/webapi/client"
	"github.com/i3odja/osbb/webapi/server"
)

func main() {
	logger := logger.NewLogger("osbb-webapi")

	logger.Infoln("Webapi service starting...")

	c, err := client.NewNotifications(os.Getenv("OSBB_NOTIFICATIONS_ADDRESS"))
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
