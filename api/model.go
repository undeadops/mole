package main

// Define our message object
type Post struct {
	Text           string `json:"text"`
	CreatedAt      string `json:"created_at"`
	RetweetCount   int    `json:"retweet_count"`
	ProfileImage   string `json:"profile_image_url"`
	ScreenName     string `json:"screen_name"`
	Name           string `json:"name"`
	FollowerCount  int    `json:"followers_count"`
	FavoritesCount int    `json:"favourites_count"`
}
