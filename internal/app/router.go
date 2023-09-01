package app

import (
	"net/http"
	"strings"
	"tictactoe-api/internal/app/handler"
)

type Router struct {
	handler.Game
}

func NewRouter(handler handler.Game) *Router {
	return &Router{handler}
}

const (
	base = `/game`
	move = `/game/move`
)

// ServeHTTP redirect request based on http verb and url
func (t *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && strings.EqualFold(base, r.URL.Path):
		t.Get(w, r)
		return
	case r.Method == http.MethodPost && strings.EqualFold(move, r.URL.Path):
		t.Move(w, r)
		return
	case r.Method == http.MethodDelete && strings.EqualFold(base, r.URL.Path):
		t.Delete(w, r)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("url not found"))
		return
	}
}
