package main

import (
  "os"
  "github.com/dghubble/go-twitter/twitter"
  "github.com/dghubble/oauth1"
  "github.com/bradfitz/slice"
  //"reflect"
)

type SelectedTweet struct {
  User string
  Text string
  CreatedAt string
  RetweetCount int
}

func getClient() (*twitter.Client) {
  config := oauth1.NewConfig(os.Getenv("TC_CONSUMER"), os.Getenv("TC_SECRET"))
  token := oauth1.NewToken(os.Getenv("TA_ACCESS"), os.Getenv("TA_SECRET"))
  httpClient := config.Client(oauth1.NoContext, token)

  return twitter.NewClient(httpClient)
}

 func getTweetsFromUser(user string) ([]SelectedTweet) {
  client := getClient()
  userTimelineParams := &twitter.UserTimelineParams{
    ScreenName: user, Count: 25 /*IncludeRetweets: false*/}
  tweets, _, _ := client.Timelines.UserTimeline(userTimelineParams)
  selectedTweets := make([]SelectedTweet, 1)

  slice.Sort(tweets[:], func(i, j int) bool {
    return tweets[i].RetweetCount > tweets[j].RetweetCount
  })

  selectedTweets[0] = SelectedTweet{
    CreatedAt: tweets[0].CreatedAt,
    Text: tweets[0].Text, 
    RetweetCount: tweets[0].RetweetCount, 
    User: user}

  return selectedTweets
}

func getTweets() ([]SelectedTweet) {  
  tweets := []SelectedTweet{}
  for i := 0; i < len(users); i++ {
    userTweets := getTweetsFromUser(users[i])
    tweets = append(tweets, userTweets...)
  }
  return tweets
}
