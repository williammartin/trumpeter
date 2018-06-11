package main

import (
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	config := oauth1.NewConfig(consumerKey(), consumerSecret())
	token := oauth1.NewToken(accessToken(), accessTokenSecret())
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	tweets, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{ScreenName: "realDonaldTrump", Count: 1})
	exitOn(err)

	fmt.Printf("\"%s\" - President Trump", mostRecentTweetText(tweets))
}

func consumerKey() string {
	return os.Getenv("CONSUMER_KEY")
}

func consumerSecret() string {
	return os.Getenv("CONSUMER_SECRET")
}

func accessToken() string {
	return os.Getenv("ACCESS_TOKEN")
}

func accessTokenSecret() string {
	return os.Getenv("ACCESS_TOKEN_SECRET")
}

func mostRecentTweetText(tweets []twitter.Tweet) string {
	return tweets[0].Text
}

func exitOn(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
