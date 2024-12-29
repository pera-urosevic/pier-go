package database

import (
	"pier/api/tvshows/types"
	"pier/storage"
)

func CreateTVShow(record types.TVShow) (types.TVShow, error) {
	db := storage.DB()

	res, err := db.Exec("INSERT INTO `tvshows` (`title`, `status`, `premiered`, `tvmaze`, `imdb`, `website`, `updated`, `episodes`, `watching`, `image`, `stream`, `runtime`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", record.Title, record.Status, record.Premiered, record.TVMaze, record.IMDB, record.Website, record.Updated, record.Episodes, record.Watching, record.Image, record.Stream, record.Runtime)
	if err != nil {
		return record, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return record, err
	}

	record.ID = id
	return record, nil
}
