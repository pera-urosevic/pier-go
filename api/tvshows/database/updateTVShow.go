package database

import (
	"pier/api/tvshows/types"
	"pier/storage"
)

func UpdateTVShow(id int64, record types.TVShow) error {
	db := storage.DB()

	_, err := db.Exec("UPDATE `tvshows` SET `title`=?, `status`=?, `premiered`=?, `tvmaze`=?, `imdb`=?, `website`=?, `updated`=?, `episodes`=?, `watching`=?, `image`=?, `stream`=?, `runtime`=? WHERE `id`=? ", record.Title, record.Status, record.Premiered, record.TVMaze, record.IMDB, record.Website, record.Updated, record.Episodes, record.Watching, record.Image, record.Stream, record.Runtime, id)
	if err != nil {
		return err
	}

	return nil
}
