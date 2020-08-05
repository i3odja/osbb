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

func NewGRPCServer(db *sql.DB, conn *Connections) *grpcServer {
	return &grpcServer{
		UnimplementedServiceServer: pb.UnimplementedServiceServer{},
		push:                       service.Push{},
		db:                         db,
		conn:                       conn,
	}
}

func (s *grpcServer) ListenAndServeGRPC(ctx context.Context, logger *logrus.Entry, addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	logger.WithField("address", addr).Infoln("GRPC server is started")

	ns := grpc.NewServer()

	pb.RegisterServiceServer(ns, s)
	if err := ns.Serve(lis); err != nil {
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
