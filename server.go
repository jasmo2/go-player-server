package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// PlayerStore server.go
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

// PlayerServer struct
type PlayerServer struct {
	store PlayerStore
	http.Handler
}

// Player Struct
type Player struct {
	Name string
	Wins int
}

// NewPlayerServer creates a PlayerServer with routing configured.
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))

	p.Handler = router

	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(p.getLeagueTable())
	w.WriteHeader((http.StatusOK))
}

func (p *PlayerServer) getLeagueTable() []Player {
	return []Player{
		{"Chris", 20},
	}
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
