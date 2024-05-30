package newsstand

import (
	"fmt"
	"os"
	"time"

	"somnusalis.org/pier/database"
	"somnusalis.org/pier/newsstand/net"
	"somnusalis.org/pier/newsstand/storage"
	"somnusalis.org/pier/notify"
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
			notify.ErrorWarn("newsstand", "skip fresh time parse", err)
			then = time.Unix(0, 0)
		}
		diff := now.Sub(then).Minutes()
		if diff < 30 {
			continue
		}

		// fetch feed
		res, err := net.Fetch(feed)
		if err != nil {
			notify.ErrorAlert("newsstand:fetch", "fetch feed "+feed.Url, err)
			continue
		}
		feed.Updated = now.Format(time.RFC3339)
		status := fmt.Sprintf("fetched %s", feed.Id)
		notify.Info("newsstand", status)

		// store articles
		storage.Articles(feed, res.Items)

		// store meta
		storage.FeedUpdate(feed)
	}
}

func check(lastRun time.Time) bool {
	key := "newsstand:reload"
	db := database.Connect()
	reload := db.Get(database.Ctx, key).Val()
	if reload != "" {
		db.Set(database.Ctx, key, "", 0)
		notification := fmt.Sprintf("reload %s", reload)
		notify.Info("newsstand", notification)
		return true
	}
	now := time.Now()
	diff := now.Sub(lastRun).Minutes()
	return diff >= 15
}

func Run() {
	if os.Getenv("RUN_NEWSSTAND") != "true" {
		return
	}

	fmt.Println("NEWSTAND")

	task()
	lastRun := time.Now()
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		if check(lastRun) {
			task()
			lastRun = time.Now()
		}
	}
}
