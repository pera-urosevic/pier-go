package database

import (
	"pier/api/tvshows/types"
	"pier/storage"
)

func GetTVShow(id int64) (types.TVShow, error) {
	var tvshow types.TVShow

	db := storage.DB()

	row := db.QueryRow("SELECT * FROM `tvshows` WHERE `id` = ?", id)
	err := row.Scan(&tvshow.ID, &tvshow.Title, &tvshow.Status, &tvshow.Premiered, &tvshow.TVMaze, &tvshow.IMDB, &tvshow.Website, &tvshow.Updated, &tvshow.Episodes, &tvshow.Watching, &tvshow.Image, &tvshow.Stream, &tvshow.Runtime)
	if err != nil {
		return tvshow, err
	}

	return tvshow, nil
}
