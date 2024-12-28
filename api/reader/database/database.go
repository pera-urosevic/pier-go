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

func UpdateFeed(name string, feed types.Feed) (types.Feed, error) {
	db := storage.DB()
	_, err := db.Exec("UPDATE `reader_feeds` SET `url` = ?, `web` = ?, `icon` = ?, `tokens` = ?, `disabled` = ?, `updated` = ?, `style` = ? WHERE `name` = ?", feed.URL, feed.Web, feed.Icon, feed.Tokens, feed.Disabled, feed.Updated, feed.Style, name)
	if err != nil {
		return feed, err
	}
	return feed, nil
}

func RemoveFeed(name string) error {
	db := storage.DB()
	_, err := db.Exec("DELETE FROM `reader_feeds` WHERE `name` = ?", name)
	if err != nil {
		return err
	}
	return nil
}

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

func DiscardFeed(name string) error {
	db := storage.DB()
	_, err := db.Exec("UPDATE `reader_articles` SET `discarded` = 1 WHERE `feed_name` = ?", name)
	if err != nil {
		return err
	}
	return nil
}

func DiscardArticle(id string) error {
	db := storage.DB()
	_, err := db.Exec("UPDATE `reader_articles` SET `discarded` = 1 WHERE `id` = ?", id)
	if err != nil {
		return err
	}
	return nil
}
