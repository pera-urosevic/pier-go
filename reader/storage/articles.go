package storage

import (
	"encoding/json"
	"fmt"
	"time"

	"pier/notify"
	"pier/reader/models"
	"pier/storage"

	"github.com/mmcdole/gofeed"
)

func Articles(feed *models.Feed, items []*gofeed.Item, threshold time.Duration) {
	db := storage.DB()

	// get db articles
	articles := map[string]*models.Article{}
	rows, err := db.Query("SELECT `id`, `content`, `discarded` FROM `reader_articles` WHERE `feed_name` = ?", feed.Name)
	if err != nil {
		notify.ErrorAlert("reader", "articles", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.Id, &article.Content, &article.Discarded); err != nil {
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

		// get article datetime
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
		// skip articles that are too old
		if dt.Before(time.Now().Add(-1 * threshold)) {
			continue
		}
		// skip articles that already exist
		id := fmt.Sprintf("%s|%s", dt, guid)
		_, exists := articles[id]
		if exists {
			delete(articles, id)
			continue
		}

		// add article to db
		db.Exec("INSERT INTO `reader_articles` (`id`, `feed_name`, `content`, `discarded`) VALUES (?, ?, ?, ?)", id, feed.Name, string(data), 0)
	}

	// delete discarded articles no longer present in feed
	for articleId, article := range articles {
		if article.Discarded {
			db.Exec("DELETE FROM `reader_articles` WHERE `id` = ?", articleId)
		}
	}
}
