package database

import (
	"pier/api/reader/types"
	"pier/storage"
)

func UpdateFeed(name string, feed types.Feed) (types.Feed, error) {
	db := storage.DB()

	_, err := db.Exec("UPDATE `reader_feeds` SET `url` = ?, `web` = ?, `icon` = ?, `tokens` = ?, `disabled` = ?, `updated` = ?, `style` = ? WHERE `name` = ?", feed.URL, feed.Web, feed.Icon, feed.Tokens, feed.Disabled, feed.Updated, feed.Style, name)
	if err != nil {
		return feed, err
	}

	return feed, nil
}
