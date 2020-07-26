package main

import (
	"context"
	"fmt"
	"os"

	"github.com/i3odja/osbb/notifications/config"
	"github.com/i3odja/osbb/notifications/controller"
	"github.com/i3odja/osbb/notifications/storage"
	"github.com/i3odja/osbb/shared/logger"
	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	logger := logger.NewLogger("osbb-notifications")

	nc, err := config.NewConfig()
	if err != nil {
		logger.WithError(err).Fatalln("could not get new config...")
	}

	dbConfig, err := nc.DBConfig(ctx)
	if err != nil {
		logger.WithError(err).Infoln("could not get db config...")
	}

	db, err := storage.ConnectToDB(dbConfig)
	if err != nil {
		logger.WithError(err).Infoln("connection to db failed!")
	}
	logger.Infoln("Connection to db successful!")

	_ = storage.NewNotifications(db)

	logger.Infoln("Starting all servers...")

	// HTTP Server Running...
	g.Go(func() error {
		err := controller.ServerAndListenHTTPServer(ctx, logger, os.Getenv("HTTP_ADDRESS"))
		if err != nil {
			return fmt.Errorf("http server failed: %w", err)
		}

		return nil
	})

	// GRPC Server Running...
	g.Go(func() error {
		err := controller.ListenAndServeGRPC(ctx, logger, db, os.Getenv("GRPC_ADDRESS"))
		if err != nil {
			return fmt.Errorf("grpc server failed: %w", err)
		}

		return nil
	})

	// WebSocket Server Running...
	g.Go(func() error {
		err := controller.ListenAndServeWebSocket(ctx, logger, os.Getenv("WEBSOCKET_ADDRESS"))
		if err != nil {
			return fmt.Errorf("websocket server failed: %w", err)
		}

		return nil
	})

	err = g.Wait()
	if err != nil {
		logger.WithError(err).Fatalln("servers failed")
	}

	logger.Infoln("All servers are terminated!")
}
