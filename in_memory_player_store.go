package main

type InMemoryPlayerStore struct {
  store map[string]int
}

func NewInMemoryPlayerStore() * InMemoryPlayerStore {
  return &InMemoryPlayerStore{store: map[string]int{}}
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
  i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
  return i.store[name]
}
