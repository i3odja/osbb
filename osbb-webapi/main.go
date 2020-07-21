package main

import (
	"context"
	"log"
	"time"

	"github.com/i3odja/osbb/webapi/client"
	"github.com/i3odja/osbb/webapi/server"
)

const address = "localhost:9999"

func main() {
	c, err := client.NewNotifications(address)
	if err != nil {
		log.Fatalf("could not create notifications client: %v", err)
	}
	defer c.Close()

	// Contact the controller and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = server.AllNotifications(ctx, c)
	if err != nil {
		log.Fatalf("could not send all notifications %v", err)
	}
}
