package main

import "sync"

type InMemoryPlayerStore struct {
  store map[string]int
  lock sync.Mutex
}

func NewInMemoryPlayerStore() * InMemoryPlayerStore {
  return &InMemoryPlayerStore{
    store: map[string]int{},
    // A mutex is used to synchronize read/write access to the map
    lock: sync.Mutex{},
  }
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
  i.lock.Lock()
  defer i.lock.Unlock()
  i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
  i.lock.Lock()
  defer i.lock.Unlock()
  return i.store[name]
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
  var league []Player
  for name, wins := range i.store {
    league = append(league, Player{name, wins})
  }

  return league
}
