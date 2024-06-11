package storage

import (
	"encoding/json"
	"fmt"
	"time"

	"pier/database"
	"pier/notify"
	"pier/reader/models"

	"github.com/mmcdole/gofeed"
)

func Articles(feed *models.Feed, items []*gofeed.Item) {
	db := database.Connect()

	articles := map[string]*models.Article{}
	rows, err := db.Query("SELECT `id`, `content` FROM `reader_articles` WHERE `feed_name` = ?", feed.Name)
	if err != nil {
		notify.ErrorAlert("reader", "articles", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.Id, &article.Content); err != nil {
			notify.ErrorAlert("reader", "articles", err)
		}
		articles[article.Id] = &article
	}

	for _, item := range items {
		data, err := json.Marshal(item)
		if err != nil {
			notify.ErrorWarn("reader", "json marshal", err)
			continue
		}

		guid := item.GUID
		var dt time.Time
		datetime := item.UpdatedParsed
		if datetime == nil {
			datetime = item.PublishedParsed
		}
		if datetime == nil {
			dt = time.Unix(feed.Updated, 0)
		} else {
			dt = *datetime
		}
		id := fmt.Sprintf("%s|%s", dt, guid)
		_, exists := articles[id]
		if exists {
			delete(articles, id)
			continue
		}

		db.Exec("INSERT INTO `reader_articles` (`id`, `feed_name`, `content`) VALUES (?, ?, ?)", id, feed.Name, string(data))
	}

	for articleId, article := range articles {
		if article.Content == "" {
			db.Exec("DELETE FROM `reader_articles` WHERE `id` = ?", articleId)
		}
	}
}
