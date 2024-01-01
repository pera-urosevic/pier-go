package newsstand

import (
	"fmt"
	"localhost/pier/newsstand/net"
	"localhost/pier/newsstand/storage"
	"os"
	"time"
)

func task() {
	feeds := storage.Feeds()
	for _, feed := range feeds {
		// skip disabled
		if feed.Disabled {
			continue
		}

		// skip fresh
		now := time.Now()
		then, err := time.Parse(time.RFC3339, feed.Updated)
		if err != nil {
			fmt.Println(err)
			then = time.Unix(0, 0)
		}
		diff := now.Sub(then).Minutes()
		if diff < 30 {
			continue
		}

		// fetch feed
		res, err := net.Fetch(feed)
		if err != nil {
			fmt.Println(err, feed.Url)
			continue
		}
		feed.Updated = now.Format(time.RFC3339)
		status := fmt.Sprintf("Fetched %s [%s]", feed.Id, feed.Updated)
		fmt.Println(status)

		// store articles
		storage.Articles(feed, res.Items)

		// store meta
		storage.FeedUpdate(feed)
	}
}

func Run() {
	if os.Getenv("RUN_NEWSSTAND") != "true" {
		return
	}

	task()
	ticker := time.NewTicker(15 * time.Minute)
	for range ticker.C {
		task()
	}
}
