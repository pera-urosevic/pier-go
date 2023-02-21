package net

import (
	"localhost/pier/newsstand/models"

	"github.com/mmcdole/gofeed"
)

func Fetch(feed *models.Feed) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	res, err := fp.ParseURL(feed.Url)
	return res, err
}
