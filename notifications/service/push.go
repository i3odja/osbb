package service

import (
	"context"
	"fmt"

	pb "github.com/i3odja/osbb/contracts/notifications"
)

type Push struct {
}

func (s *Push) PersonalNotification(ctx context.Context, userID string, notification []*pb.Notification) (string, error) {
	return fmt.Sprintf("my ID is %s and I'm Personal Notification!", userID), nil
}

func (s *Push) EverybodyNotification(ctx context.Context, notification []*pb.Notification) (string, error) {
	return "I'm EverybodyNotification!", nil
}

func (s *Push) PersonalTestNotification(ctx context.Context, notification []*pb.Notification) (string, error) {
	return fmt.Sprintf("Title: %s and I'm Personal Test Notification!", notification[0].Title), nil
}
