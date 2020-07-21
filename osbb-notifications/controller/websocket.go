package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func echo(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

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

func hello(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

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

func ListenAndServeWebSocket(ctx context.Context, addr string) error {
	mainRoute := mux.NewRouter()
	ws := mainRoute.PathPrefix("/ws").Subrouter()
	ws.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/websockets.html")
	})
	ws.HandleFunc("/echo", echo)
	ws.HandleFunc("/hello", hello)

	fmt.Printf(" + [Websocket server listening... at%s]\n", addr)

	err := http.ListenAndServe(addr, mainRoute)
	if err != nil {
		return fmt.Errorf("failed to serve web socket: %w", err)
	}

	return nil
}
