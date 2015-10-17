package handlers

import "net/http"

type Handler interface {
	Init() error
	Reply(w http.ResponseWriter, r *http.Request)
}
