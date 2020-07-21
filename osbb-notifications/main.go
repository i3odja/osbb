package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/i3odja/osbb/notifications/controller"
	"github.com/i3odja/osbb/shared/logger"
	"golang.org/x/sync/errgroup"
)

var httpAddress = flag.String("http_addr", ":8189", "HTTP service address")
var grpcAddress = flag.String("grpc_addr", ":9999", "GRPC service address")
var wsAddress = flag.String("ws_addr", ":9090", "WebSocket service address")

func main() {
	logger := logger.NewLogger("osbb-notifications")
	flag.Parse()

	g, ctx := errgroup.WithContext(context.Background())

	logger.Infoln("Starting all servers...")

	// HTTP Server Running...
	g.Go(func() error {
		err := controller.ServerAndListenHTTPServer(ctx, logger, *httpAddress)
		if err != nil {
			return fmt.Errorf("http server failed: %w", err)
		}

		return nil
	})

	// GRPC Server Running...
	g.Go(func() error {
		err := controller.ListenAndServeGRPC(ctx, logger, *grpcAddress)
		if err != nil {
			return fmt.Errorf("grpc server failed: %w", err)
		}

		return nil
	})

	// WebSocket Server Running...
	g.Go(func() error {
		err := controller.ListenAndServeWebSocket(ctx, logger, *wsAddress)
		if err != nil {
			return fmt.Errorf("websocket server failed: %w", err)
		}

		return nil
	})

	err := g.Wait()
	if err != nil {
		logger.WithError(err).Fatalln("servers failed")
	}

	logger.Infoln("All servers are terminated!")
}
