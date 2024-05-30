package storage

import (
	"fmt"

	"somnusalis.org/pier/database"
	"somnusalis.org/pier/newsstand/models"
)

func FeedUpdate(feed *models.Feed) {
	db := database.Connect()
	key := fmt.Sprintf("newsstand:%s:feed", feed.Id)
	db.HSet(database.Ctx, key, "updated", feed.Updated)
}
