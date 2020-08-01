package server

import (
	"context"
	"fmt"

	pb "github.com/i3odja/osbb/contracts/notifications"
	"github.com/i3odja/osbb/webapi/client"
	"github.com/sirupsen/logrus"
)

func AllNotifications(ctx context.Context, logger *logrus.Entry, c *client.Notifications) error {
	r, err := c.SendNotification(ctx, &pb.SendRequest{UserId: "0673419017", Notification: nil})
	if err != nil {
		return fmt.Errorf("could not send notification: %w", err)
	}
	logger.WithField("response", r.GetSResponse()).Debugln("Send notification")

	rb, err := c.BroadcastNotification(ctx, &pb.BroadcastRequest{Notification: nil})
	if err != nil {
		return fmt.Errorf("could not broadcast notification: %w", err)
	}
	logger.WithField("response", rb.GetBResponse()).Debugln("Broadcast notification")

	rm, err := c.MyNotification(ctx, &pb.MyRequest{
		Notification: []*pb.Notification{
			{
				Title: "My testing",
			},
		},
	})
	if err != nil {
		return fmt.Errorf("could not my notification: %w", err)
	}
	logger.WithField("response", rm.GetMResponse()).Debugln("My notifications")

	rs, err := c.SendNotification(ctx, &pb.SendRequest{
		UserId: "0731674016",
	})
	if err != nil {
		return fmt.Errorf("could not send notification: %w", err)
	}
	logger.WithField("response", rs.GetSResponse()).Debugln("Send notifications with uid")

	return nil
}
