package net

import (
	"pier/reader/models"

	"github.com/mmcdole/gofeed"
)

func Fetch(feed *models.Feed) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	fp.UserAgent = "Pier 2.0"
	res, err := fp.ParseURL(feed.Url)
	return res, err
}
