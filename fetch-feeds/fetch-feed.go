package main

import (
	"github.com/araddon/dateparse"
	"github.com/mmcdole/gofeed"
	"time"
)

func FetchFeed(url string) *gofeed.Feed {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(url)
	wa := time.Now().AddDate(0, 0, -7)
	is := make([]*gofeed.Item, 0)

	for _, v := range feed.Items {
		t, err := dateparse.ParseAny(v.Published)

		if err != nil {
			panic("Bad date")
		}

		if t.After(wa) {
			is = append(is, v)
		}
	}

	feed.Items = is

	return feed
}
