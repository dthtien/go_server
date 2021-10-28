package main

import (
  "log"
  "net/http"
)

func main() {
  server := &PlayerServer{}
  log.Fatal(http.ListenAndServe(":5000", server))
}

