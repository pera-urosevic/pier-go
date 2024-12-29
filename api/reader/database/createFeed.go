package database

import (
	"pier/api/reader/types"
	"pier/storage"
)

func CreateFeed(name string) (types.Feed, error) {
	var feed = types.Feed{
		Name:     name,
		URL:      "",
		Web:      "",
		Icon:     "",
		Tokens:   "{}",
		Disabled: true,
		Updated:  0,
		Style:    "",
	}

	db := storage.DB()

	_, err := db.Exec("INSERT INTO `reader_feeds` (`name`, `url`, `web`, `icon`, `tokens`, `style`, `disabled`, `updated`) VALUES(?, ?, ?, ?, ?, ?, ?, ?)", feed.Name, feed.URL, feed.Web, feed.Icon, feed.Tokens, feed.Style, feed.Disabled, feed.Updated)
	if err != nil {
		return feed, err
	}

	return feed, nil
}
