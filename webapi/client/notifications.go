package client

import (
	"fmt"

	pb "github.com/i3odja/osbb/contracts/notifications"
	"google.golang.org/grpc"
)

type Notifications struct {
	pb.ServiceClient
	close func() error
}

func NewNotifications(address string) (*Notifications, error) {
	// Set up a connection to the controller.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("did not connect: %w", err)
	}

	return &Notifications{
		ServiceClient: pb.NewServiceClient(conn),
		close:         conn.Close,
	}, nil
}

func (n *Notifications) Close() error {
	return n.close()
}
