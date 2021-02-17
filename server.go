package main

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerStore server.go
type PlayerStore interface {
	GetPlayerScore(name string) int
}

// PlayerServer struct
type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, p.store.GetPlayerScore(player))
}