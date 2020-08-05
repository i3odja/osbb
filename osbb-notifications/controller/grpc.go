package controller

import (
	"context"
	"database/sql"
	"fmt"
	"net"

	pb "github.com/i3odja/osbb/contracts/notifications"
	"github.com/i3odja/osbb/notifications/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type grpcServer struct {
	pb.UnimplementedServiceServer
	push service.Push
	db   *sql.DB
	conn *Connections
}

func ListenAndServeGRPC(ctx context.Context, logger *logrus.Entry, n *sql.DB, conn *Connections, addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	logger.WithField("address", addr).Infoln("GRPC server is started")

	s := grpc.NewServer()

	pb.RegisterServiceServer(s, &grpcServer{db: n, conn: conn})
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

func (s *grpcServer) SendNotification(ctx context.Context, in *pb.SendRequest) (*pb.SendResponse, error) {
	id, err := s.push.PersonalNotification(ctx, in.UserId, in.Notification)
	if err != nil {
		return &pb.SendResponse{SResponse: err.Error()}, nil
	}

	n := service.NewNotifications(s.db)
	err = n.Add(id)
	if err != nil {
		return nil, fmt.Errorf("add error: %w", err)
	}

	go s.conn.broadcastMessage(id)

	return &pb.SendResponse{SResponse: id}, nil
}

func (s *grpcServer) BroadcastNotification(ctx context.Context, in *pb.BroadcastRequest) (*pb.BroadcastResponse, error) {
	id, err := s.push.EverybodyNotification(ctx, in.Notification)
	if err != nil {
		return &pb.BroadcastResponse{BResponse: err.Error()}, nil
	}

	n := service.NewNotifications(s.db)
	err = n.Add(id)
	if err != nil {
		return nil, fmt.Errorf("add error: %w", err)
	}

	go s.conn.broadcastMessage(id)

	return &pb.BroadcastResponse{BResponse: id}, nil
}

func (s *grpcServer) MyNotification(ctx context.Context, in *pb.MyRequest) (*pb.MyResponse, error) {
	title, err := s.push.PersonalTestNotification(ctx, in.Notification)
	if err != nil {
		return &pb.MyResponse{MResponse: err.Error()}, nil
	}

	n := service.NewNotifications(s.db)
	err = n.Add(title)
	if err != nil {
		return nil, fmt.Errorf("add error: %w", err)
	}

	go s.conn.broadcastMessage(title)

	return &pb.MyResponse{MResponse: title}, nil
}
