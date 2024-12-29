package database

import (
	"pier/api/tvshows/types"
	"pier/storage"
)

func GetTVShows() ([]types.TVShow, error) {
	db := storage.DB()

	rows, err := db.Query("SELECT * FROM `tvshows`")
	if err != nil {
		return nil, err
	}

	var tvshows []types.TVShow
	for rows.Next() {
		var tvshow types.TVShow
		err := rows.Scan(&tvshow.ID, &tvshow.Title, &tvshow.Status, &tvshow.Premiered, &tvshow.TVMaze, &tvshow.IMDB, &tvshow.Website, &tvshow.Updated, &tvshow.Episodes, &tvshow.Watching, &tvshow.Image, &tvshow.Stream, &tvshow.Runtime)
		if err != nil {
			return nil, err
		}

		tvshows = append(tvshows, tvshow)
	}

	return tvshows, nil
}
