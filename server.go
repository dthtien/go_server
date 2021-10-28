package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
  GetPlayerScore(name string) int
}

type PlayerServer struct {
  store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
  case http.MethodPost:
    p.processWin(w)
  case http.MethodGet:
    p.showScore(w, r)
  }
}

func (p *PlayerServer) processWin(w http.ResponseWriter) {
  w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
  player := strings.TrimPrefix(r.URL.Path, "/players/")
  score := p.store.GetPlayerScore(player)

  if score == 0 {
    w.WriteHeader(http.StatusNotFound)
  }

  fmt.Fprint(w, score)
}

func GetPlayerScore(name string) string{
  if name == "Pepper" {
    return "20"
  }

  if name == "Floyd" {
    return "10"
  }

  return ""
}
