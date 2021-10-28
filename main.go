package main

import (
  "log"
  "net/http"
)

type InmemoryPlayerStore struct{}

func(i *InmemoryPlayerStore) GetPlayerScore(name string) int {
  return 123
}

func main() {
  server := &PlayerServer{&InmemoryPlayerStore{}}
  log.Fatal(http.ListenAndServe(":5000", server))
}

