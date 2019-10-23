package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	command *string
)

// WordData is data that will help with training a model that detects names.
type WordData struct {
	Word       string
	Occurances int
	// Vector     word2vec.Vector
	IsName bool
}

// Tweet object for holding raw form of csv data.
type Tweet struct {
	ID            int
	Lang          string
	Date          int64
	Source        string
	Length        int
	OrgTweet      string
	Tweets        string
	Likes         int
	Retweets      int
	Hashtag       string
	Usermention   string
	UserMentionID string
	Name          string
	Place         string
	Followers     int
	Friends       int
}

func init() {
	command = flag.String("action", "train", "This is the method that the program in performing.")
}
func main() {
	flag.Parse()
	fmt.Println(*command)
	file := "../datasets/FIFA.csv"
	switch *command {
	case "plot":
		err := UnderstandData(file)
		log.Println(err)
	case "train":
		err := DecisionTree(("data/words.csv"))
		log.Println((err))
	}
}
