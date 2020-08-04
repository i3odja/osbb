package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/i3odja/osbb/notifications/service"
	"github.com/sirupsen/logrus"
)

type HTTP struct {
	notification *service.Notifications
}

func NewHTTP(notification *service.Notifications) *HTTP {
	return &HTTP{
		notification: notification,
	}
}

func (h *HTTP) Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I'm testing endpoint..."))
}

func (h *HTTP) GetID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]

	gotMessage, err := h.notification.Get(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	responseMessage := fmt.Sprintf("My ID: %s\n\tmessage: %s", id, gotMessage)

	w.Write([]byte(responseMessage))
}

func (h *HTTP) ServerAndListenHTTPServer(ctx context.Context, logger *logrus.Entry, addr string) error {
	rr := mux.NewRouter()
	rr.HandleFunc("/test", h.Test)
	rr.HandleFunc("/test/{ID}", h.GetID)
	http.Handle("/", rr)

	logger.WithField("address", addr).Infoln("HTTP server is started")

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		return fmt.Errorf("failed to serve http server: %w", err)
	}

	return nil
}
