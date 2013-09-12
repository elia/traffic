package main

import (
  "net/http"
  "github.com/pilu/traffic"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
  panic("this is a test error!")
}

func main() {
  router := traffic.New()

  // Routes
  router.Get("/", rootHandler)

  http.Handle("/", router)
  http.ListenAndServe(":7000", nil)
}
