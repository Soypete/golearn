package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/kniren/gota/dataframe"
)

var (
	command *string
)

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
		err := plotData(file)
		log.Println(err)
	default:

	}
}

func plotData(file string) error {
	tweets, err := getData(file)
	if err != nil {
		return fmt.Errorf("error getting tweets %v", err)
	}
	df := dataframe.LoadStructs(tweets)
	fmt.Println(df)
	fmt.Println(df.Select([]string{"Tweets"}))
	return nil
}

func getData(file string) ([]Tweet, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}
	var tweets []Tweet
	for i, line := range lines {
		if i == 0 {
			continue
		}
		// fmt.Println(line)
		id, _ := strconv.Atoi(line[0])
		layout := "2006-01-02"
		d, _ := time.Parse(layout, line[2])
		l, _ := strconv.Atoi(line[4])
		likes, _ := strconv.Atoi(line[7])
		r, _ := strconv.Atoi(line[8])
		f, _ := strconv.Atoi(line[14])
		friends, _ := strconv.Atoi(line[14])

		t := Tweet{
			ID:            id,
			Lang:          line[1],
			Date:          d.Unix(),
			Source:        line[3],
			Length:        l,
			OrgTweet:      line[5],
			Tweets:        line[6],
			Likes:         likes,
			Retweets:      r,
			Hashtag:       line[9],
			Usermention:   line[10],
			UserMentionID: line[11],
			Name:          line[12],
			Place:         line[13],
			Followers:     f,
			Friends:       friends,
		}
		tweets = append(tweets, t)
	}
	fmt.Println(len(tweets))
	return tweets, err
}
