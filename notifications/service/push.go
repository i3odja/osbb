package service

import (
	"context"

	pb "github.com/i3odja/osbb/contracts/notifications"
)

type Push struct {
}

func (s *Push) PersonalNotification(ctx context.Context, userID string, notification []*pb.Notification) (string, error) {
	return "I'm Personal Notification!", nil
}

func (s *Push) EverybodyNotification(ctx context.Context, notification []*pb.Notification) (string, error) {
	return "I'm EverybodyNotification!", nil
}
