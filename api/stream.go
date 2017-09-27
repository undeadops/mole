package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

var (
	wait    = 1
	maxWait = 300
)

func min(min, max int) int {
	if min > max {
		return max
	}
	return min
}

func readStream() {
	consumerKey := os.Getenv("CONSUMER-KEY")
	consumerSec := os.Getenv("CONSUMER-SECRET")
	accessToken := os.Getenv("ACCESS-TOKEN")
	accessSec := os.Getenv("ACCESS-SECRET")
	track := os.Getenv("TRACK")

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSec)
	client := anaconda.NewTwitterApi(accessToken, accessSec)

	client.SetLogger(anaconda.BasicLogger)

	v := url.Values{}
	v.Set("track", track)
	v.Set("language", "en")
	fmt.Println("Starting Public Stream Filter...")
	s := client.PublicStreamFilter(v)

	for {
		item := <-s.C
		switch status := item.(type) {
		case anaconda.Tweet:
			p := Post{
				Text:           status.Text,
				CreatedAt:      status.CreatedAt,
				RetweetCount:   status.RetweetCount,
				ProfileImage:   status.User.ProfileImageURL,
				ScreenName:     status.User.ScreenName,
				Name:           status.User.Name,
				FollowerCount:  status.User.FollowersCount,
				FavoritesCount: status.FavoriteCount,
			}
			broadcast <- p
		default:
		}
	}

}
