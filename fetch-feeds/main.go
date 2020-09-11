package main

import (
	"fmt"
)

type FeedSpec struct {
	URL string
}

func main() {
	feeds := []FeedSpec{
		FeedSpec{"https://4health.se/feed"},
		FeedSpec{"https://www.hemnet.se/mitt_hemnet/sparade_sokningar/26325388.xml"},
	}

	for _, v := range feeds {
		f := FetchFeed(v.URL)
		fmt.Printf("Sending %s email\n", f.Title)
		SendEmail(f)
	}
}
