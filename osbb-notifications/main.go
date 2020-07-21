package main

import (
	"context"
	"flag"
	"sync"

	"github.com/i3odja/osbb/shared/logger"

	"github.com/i3odja/osbb/notifications/controller"
)

var httpAddress = flag.String("http_addr", ":8189", "HTTP service address")
var grpcAddress = flag.String("grpc_addr", ":9999", "GRPC service address")
var wsAddress = flag.String("ws_addr", ":9090", "WebSocket service address")

func main() {
	logger := logger.NewLogger("osbb-notifications")
	flag.Parse()

	var wg sync.WaitGroup

	wg.Add(3)

	logger.Infoln("Starting all servers...")

	// HTTP Server Running...
	go func() {
		defer wg.Done()

		err := controller.ServerAndListenHTTPServer(context.TODO(), logger, *httpAddress)
		if err != nil {
			logger.WithError(err).Fatalln("HTTP Server failed")
		}
	}()

	// GRPC Server Running...
	go func() {
		defer wg.Done()

		err := controller.ListenAndServeGRPC(context.TODO(), logger, *grpcAddress)
		if err != nil {
			logger.WithError(err).Fatalln("GRPC Server failed")
		}
	}()

	// WebSocket Server Running...
	go func() {
		defer wg.Done()

		err := controller.ListenAndServeWebSocket(context.TODO(), logger, *wsAddress)
		if err != nil {
			logger.WithError(err).Fatalln("WebSocket Server failed")
		}
	}()

	wg.Wait()

	logger.Infoln("All servers are terminated!")
}
