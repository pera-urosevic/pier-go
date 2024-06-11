package reader

import (
	"fmt"
	"os"
	"time"

	"pier/notify"
	"pier/reader/net"
	"pier/reader/storage"
)

func task(forced bool) {
	feeds := storage.Feeds()
	for _, feed := range feeds {
		// skip disabled
		if feed.Disabled {
			continue
		}

		now := time.Now()
		if !forced {
			// skip fresh
			then := time.Unix(feed.Updated, 0)
			diff := now.Sub(then).Minutes()
			if diff < 30 {
				continue
			}
		}

		// fetch feed
		res, err := net.Fetch(feed)
		if err != nil {
			notify.ErrorAlert("reader:fetch", "fetch feed "+feed.Url, err)
			continue
		}
		feed.Updated = now.Unix()
		status := fmt.Sprintf("fetched %s", feed.Name)
		notify.Info("reader", status)

		// store articles
		storage.Articles(feed, res.Items)

		// store meta
		storage.FeedUpdate(feed)
	}
}

func check(lastRun time.Time) (bool, bool) {
	reload := storage.Reload()
	if reload {
		notify.Info("reader", "reload")
		return true, true
	}
	now := time.Now()
	diff := now.Sub(lastRun).Minutes()
	return diff >= 15, false
}

func Run() {
	if os.Getenv("RUN_READER") != "true" {
		return
	}

	fmt.Println("READER")

	task(false)
	lastRun := time.Now()
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		run, force := check(lastRun)
		if run {
			task(force)
			lastRun = time.Now()
		}
	}
}
