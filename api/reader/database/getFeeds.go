package database

import (
	"pier/api/reader/types"
	"pier/storage"
)

func GetFeeds() ([]types.Feed, error) {
	var feeds = []types.Feed{}

	db := storage.DB()

	rows, err := db.Query("SELECT * FROM `reader_feeds`")
	if err != nil {
		return feeds, err
	}

	for rows.Next() {
		var feed types.Feed

		err := rows.Scan(&feed.Name, &feed.URL, &feed.Web, &feed.Icon, &feed.Tokens, &feed.Disabled, &feed.Updated, &feed.Style)
		if err != nil {
			return feeds, err
		}

		feeds = append(feeds, feed)
	}

	return feeds, nil
}
