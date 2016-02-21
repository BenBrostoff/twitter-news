package main

import (
  "log"
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
)

func Tweets(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(getTweets())
}

func main() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/tweets", Tweets)

  log.Fatal(http.ListenAndServe(":8080", router))
}
