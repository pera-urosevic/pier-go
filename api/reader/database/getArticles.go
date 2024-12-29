package database

import (
	"pier/api/reader/types"
	"pier/storage"
)

func GetArticles() ([]types.Article, error) {
	var articles = []types.Article{}

	db := storage.DB()
	rows, err := db.Query("SELECT * FROM `reader_articles` WHERE `discarded` = 0")
	if err != nil {
		return articles, err
	}

	for rows.Next() {
		var article types.Article

		err := rows.Scan(&article.ID, &article.Content, &article.FeedName, &article.Discarded)
		if err != nil {
			return articles, err
		}

		articles = append(articles, article)
	}

	return articles, nil
}
