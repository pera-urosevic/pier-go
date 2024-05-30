package storage

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pera-urosevic/pier/database"
	"github.com/pera-urosevic/pier/newsstand/models"
	"github.com/pera-urosevic/pier/notify"

	"github.com/mmcdole/gofeed"
)

func Articles(feed *models.Feed, items []*gofeed.Item) {
	db := database.Connect()
	key := fmt.Sprintf("newsstand:%s:articles", feed.Id)

	articles, err := db.HGetAll(database.Ctx, key).Result()
	if err != nil {
		notify.ErrorInfo("newsstand", "get articles", err)
	}

	for _, item := range items {
		data, err := json.Marshal(item)
		if err != nil {
			notify.ErrorWarn("newsstand", "json marshal", err)
			continue
		}

		id := item.GUID
		var dt string
		datetime := item.UpdatedParsed
		if datetime == nil {
			datetime = item.PublishedParsed
		}
		if datetime == nil {
			dt = feed.Updated
		} else {
			dt = datetime.Format(time.RFC3339)
		}
		field := fmt.Sprintf("%s|%s", dt, id)
		_, exists := articles[field]
		if exists {
			delete(articles, field)
			continue
		}

		db.HSet(database.Ctx, key, field, string(data))
	}

	for articleField, articleValue := range articles {
		if articleValue == "" {
			db.HDel(database.Ctx, key, articleField)
		}
	}
}
