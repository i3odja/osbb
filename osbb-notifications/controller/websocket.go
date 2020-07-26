package controller

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// clients includes all active connections
var clients = make(map[*websocket.Conn]struct{})
var mu sync.Mutex
var broadcast = make(chan []byte)

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// delete client connection after closed
	defer delete(clients, conn)

	// Add new client connections to map
	mu.Lock()
	clients[conn] = struct{}{}
	mu.Unlock()

	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Print the message to the console
		fmt.Printf("%s sent to server: %s\n", conn.RemoteAddr(), string(msg))

		// Write message back to browser
		if err = conn.WriteMessage(msgType, msg); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func ListenAndServeWebSocket(ctx context.Context, logger *logrus.Entry, addr string) error {
	r := mux.NewRouter()
	r.HandleFunc("/ws", websocketHandler)

	logger.WithField("address", addr).Infoln("Websocket server is started")

	go countOfClients()

	err := http.ListenAndServe(addr, r)
	if err != nil {
		return fmt.Errorf("failed to serve web socket: %w", err)
	}

	return nil
}

// countOfClients() shows information about active connections with interval equal to 5 seconds
func countOfClients() {
	for {
		time.Sleep(5 * time.Second)
		fmt.Printf("At the moment, are connected %d clients\n", len(clients))
	}
}

// broadcastMessage() send message to all active clients
func BroadcastMessage(text string) {
	for connection, _ := range clients {
		if err := connection.WriteMessage(1, []byte(text)); err != nil {
			return
		}
	}
}
