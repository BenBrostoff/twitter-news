package main

import (
  "log"
  "strconv"
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
)

func Tweets(w http.ResponseWriter, r *http.Request) { 
  tweetHistoryStr := r.FormValue("tweet_history")
  tweetHistoryNum, _ := strconv.Atoi(tweetHistoryStr)
  if tweetHistoryNum == 0 {
    tweetHistoryNum, _ = strconv.Atoi(getDefaultTweetNum())
  }

  json.NewEncoder(w).Encode(getTweets(tweetHistoryNum))
}

func main() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/tweets", Tweets)

  log.Fatal(http.ListenAndServe(":8080", router))
}
