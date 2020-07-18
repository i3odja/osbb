package main

import (
	"context"
	"log"
	"time"

	pb "github.com/i3odja/osbb/contracts/notifications"
	"github.com/i3odja/osbb/webapi/client"
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
	r, err := c.SendNotification(ctx, &pb.SendRequest{UserId: "0673419017", Notification: nil})
	if err != nil {
		log.Fatalf("could not send notification: %v", err)
	}
	log.Printf("Sending...: %s", r.GetSResponse())

	rb, err := c.BroadcastNotification(ctx, &pb.BroadcastRequest{Notification: nil})
	if err != nil {
		log.Fatalf("could not broadcast notification: %v", err)
	}
	log.Printf("Broadcasting...: %s", rb.GetBResponse())

	rm, err := c.MyNotification(ctx, &pb.MyRequest{Notification: nil})
	if err != nil {
		log.Fatalf("could not my notification: %v", err)
	}
	log.Printf("my...: %s", rm.GetMResponse())
}
