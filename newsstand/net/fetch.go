package net

import (
	"github.com/pera-urosevic/pier/newsstand/models"

	"github.com/mmcdole/gofeed"
)

func Fetch(feed *models.Feed) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	fp.UserAgent = "Pier 1.0"
	res, err := fp.ParseURL(feed.Url)
	return res, err
}
