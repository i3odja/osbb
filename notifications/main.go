package main

import (
	"context"
	"log"

	"github.com/i3odja/osbb/notifications/controller"
)

func main() {
	addr := ":9999"

	err := controller.ListenAndServeGRPC(context.TODO(), addr)
	if err != nil {
		log.Fatalf("GRPC: %v", err)
	}
}
