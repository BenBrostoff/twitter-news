package main
import (
  "os"
  "fmt"
  "github.com/dghubble/go-twitter/twitter"
  "github.com/dghubble/oauth1"
  "github.com/bradfitz/slice"
)

func getClient() (*twitter.Client){
  config := oauth1.NewConfig(os.Getenv("TC_CONSUMER"), os.Getenv("TC_SECRET"))
  token := oauth1.NewToken(os.Getenv("TA_ACCESS"), os.Getenv("TA_SECRET"))
  httpClient := config.Client(oauth1.NoContext, token)

  return twitter.NewClient(httpClient)
}

 func getTweets() ([]twitter.Tweet) {
  client := getClient()
  tweets, _, _ := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{})

  slice.Sort(tweets[:], func(i, j int) bool {
    return tweets[i].RetweetCount > tweets[j].RetweetCount
  })

  return tweets
}

func main() {
  for _, tweet := range getTweets() {
      fmt.Printf(tweet.Text)
      fmt.Printf("\n")
  }
}
