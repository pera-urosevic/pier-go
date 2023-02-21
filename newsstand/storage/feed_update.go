package storage

import (
	"fmt"
	"localhost/pier/database"
	"localhost/pier/newsstand/models"
)

func FeedUpdate(feed *models.Feed) {
	db := database.Connect()
	key := fmt.Sprintf("newsstand:%s:feed", feed.Id)
	err := db.HSet(database.Ctx, key, "updated", feed.Updated)
	if err != nil {
		fmt.Println(err)
	}
}
