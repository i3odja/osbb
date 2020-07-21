package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
	if err != nil {
		fmt.Printf("websocketHandler: %v", err)
	}

	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// Print the message to the console
		fmt.Printf("%s sent to server: %s\n", conn.RemoteAddr(), string(msg))

		// Write message back to browser
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}

func ListenAndServeWebSocket(ctx context.Context, logger *logrus.Entry, addr string) error {
	r := mux.NewRouter()
	r.HandleFunc("/ws", websocketHandler)

	logger.WithField("address", addr).Infoln("Websocket server is started")

	err := http.ListenAndServe(addr, r)
	if err != nil {
		return fmt.Errorf("failed to serve web socket: %w", err)
	}

	return nil
}
