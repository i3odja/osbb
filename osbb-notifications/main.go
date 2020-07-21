package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/i3odja/osbb/notifications/controller"
)

var httpAddress = flag.String("http_addr", ":8189", "HTTP service address")
var grpcAddress = flag.String("grpc_addr", ":9999", "GRPC service address")
var wsAddress = flag.String("ws_addr", ":9090", "WebSocket service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	var wg sync.WaitGroup

	wg.Add(3)

	fmt.Println("[Starting all servers...]")

	// HTTP Server Running...
	go func() {
		defer wg.Done()

		err := controller.ServerAndListenHTTPServer(context.TODO(), *httpAddress)
		if err != nil {
			log.Fatalf("HTTP Server: %v", err)
		}
	}()

	// GRPC Server Running...
	go func() {
		defer wg.Done()

		err := controller.ListenAndServeGRPC(context.TODO(), *grpcAddress)
		if err != nil {
			log.Fatalf("GRPC Server: %v", err)
		}
	}()

	// WebSocket Server Running...
	go func() {
		defer wg.Done()

		err := controller.ListenAndServeWebSocket(context.TODO(), *wsAddress)
		if err != nil {
			log.Fatalf("WebSocket Server: %v", err)
		}
	}()

	wg.Wait()

	fmt.Println("All servers have terminated!")
}
