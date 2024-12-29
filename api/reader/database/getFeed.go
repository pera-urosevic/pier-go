package database

import (
	"pier/api/reader/types"
	"pier/storage"
)

func GetFeed(name string) (types.Feed, error) {
	var feed = types.Feed{}

	db := storage.DB()

	row := db.QueryRow("SELECT * FROM `reader_feeds` WHERE `name` = ?", name)
	err := row.Scan(&feed.Name, &feed.URL, &feed.Web, &feed.Icon, &feed.Tokens, &feed.Disabled, &feed.Updated, &feed.Style)
	if err != nil {
		return feed, err
	}

	return feed, nil
}
