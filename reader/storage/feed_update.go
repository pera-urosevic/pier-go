package storage

import (
	"pier/reader/models"
	"pier/storage"
)

func FeedUpdate(feed *models.Feed) {
	db := storage.DB()
	db.Exec("UPDATE `reader_feeds` SET `updated` = ? WHERE `name` = ?", feed.Updated, feed.Name)
}
