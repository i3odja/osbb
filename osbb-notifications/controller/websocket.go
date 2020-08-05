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

type Connections struct {
	conn    *websocket.Conn
	clients map[*websocket.Conn]struct{}
	mu      sync.Mutex
	addr    string
}

func NewConnections(addr string) *Connections {
	return &Connections{
		conn:    nil,
		clients: make(map[*websocket.Conn]struct{}),
		mu:      sync.Mutex{},
		addr:    addr,
	}
}

func (c *Connections) websocketHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	c.conn, err = upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// delete client connection after closed
	defer c.deleteConnection()

	c.addConnection()

	for {
		// Read message from browser
		msgType, msg, err := c.conn.ReadMessage()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Print the message to the console
		fmt.Printf("%s sent to server: %s\n", c.conn.RemoteAddr(), string(msg))

		// Write message back to browser
		if err = c.conn.WriteMessage(msgType, msg); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// deleteConnection() removes client connection after closed
func (c *Connections) deleteConnection() {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.clients, c.conn)
}

// addConnection() adds new client connections to map
func (c *Connections) addConnection() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.clients[c.conn] = struct{}{}
}

func (c *Connections) ListenAndServeWebSocket(ctx context.Context, logger *logrus.Entry) error {
	r := mux.NewRouter()
	r.HandleFunc("/ws", c.websocketHandler)

	logger.WithField("address", c.addr).Infoln("Websocket server is started")

	go c.countOfClients()

	err := http.ListenAndServe(c.addr, r)
	if err != nil {
		return fmt.Errorf("failed to serve web socket: %w", err)
	}

	return nil
}

// countOfClients() shows information about active connections with interval equal to 5 seconds
func (c *Connections) countOfClients() {
	for {
		time.Sleep(5 * time.Second)
		fmt.Printf("At the moment, are connected %d clients\n", len(c.clients))
	}
}

// broadcastMessage() send message to all active clients
func (c *Connections) broadcastMessage(text string) {
	for connection, _ := range c.clients {
		if err := connection.WriteMessage(1, []byte(text)); err != nil {
			return
		}
	}
}
