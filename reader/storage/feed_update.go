package storage

import (
	"pier/database"
	"pier/reader/models"
)

func FeedUpdate(feed *models.Feed) {
	db := database.Connect()
	db.Exec("UPDATE `reader_feeds` SET `updated` = ? WHERE `name` = ?", feed.Updated, feed.Name)
}
