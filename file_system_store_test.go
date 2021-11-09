package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
  database, cleanDatabase := createTempFile(t, `[
  { "Name": "Cleo", "Wins": 10 },
  { "name": "Chris", "Wins": 33 }]`)
  defer cleanDatabase()
  store := NewFileSystemPlayerStore(database)

  t.Run("league from a reader", func(t *testing.T) {
    got := store.GetLeague()
    want := []Player{
      {"Cleo", 10},
      {"Chris", 33},
    }

    assertLeague(t, got, want)
  })

  t.Run("get player score", func(t *testing.T) {
    got := store.GetPlayerScore("Chris")
    want := 33

    assertScoreEquals(t, got, want)
  })

  t.Run("store wins for existing players", func(t *testing.T) {
    store.RecordWin("Chris")
    got := store.GetPlayerScore("Chris")
    want := 34
    assertScoreEquals(t, got, want)
  })

  t.Run("store wins for new players", func(t *testing.T){
    store.RecordWin("Pepper")
    got := store.GetPlayerScore("Pepper")
    want := 1
    assertScoreEquals(t, got, want)
  })
}

func assertScoreEquals(t *testing.T, got, want int) {
  t.Helper()

  if got != want {
    t.Errorf("got %d want %d", got, want)
  }
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()){
  t.Helper()

  tmpFile, err := ioutil.TempFile("", "db")

  if err != nil {
    t.Fatalf("could not create temp file %v", err)
  }

  tmpFile.Write([]byte(initialData))

  removeFile := func() {
    tmpFile.Close()
    os.Remove(tmpFile.Name())
  }

  return tmpFile, removeFile
}
