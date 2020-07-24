package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I'm testing endpoint..."))
}

func ServerAndListenHTTPServer(ctx context.Context, logger *logrus.Entry, addr string) error {
	rr := mux.NewRouter()
	rr.HandleFunc("/test", Test)
	http.Handle("/", rr)

	logger.WithField("address", addr).Infoln("HTTP server is started")

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		return fmt.Errorf("failed to serve http server: %w", err)
	}

	return nil
}
