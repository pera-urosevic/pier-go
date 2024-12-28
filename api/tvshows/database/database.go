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

func UpdateTVShow(id int64, record types.TVShow) error {
	db := storage.DB()
	_, err := db.Exec("UPDATE `tvshows` SET `title`=?, `status`=?, `premiered`=?, `tvmaze`=?, `imdb`=?, `website`=?, `updated`=?, `episodes`=?, `watching`=?, `image`=?, `stream`=?, `runtime`=? WHERE `id`=? ", record.Title, record.Status, record.Premiered, record.TVMaze, record.IMDB, record.Website, record.Updated, record.Episodes, record.Watching, record.Image, record.Stream, record.Runtime, id)
	if err != nil {
		return err
	}
	return nil
}

func RemoveTVShow(id int64) error {
	db := storage.DB()
	_, err := db.Exec("DELETE FROM `tvshows` WHERE `id` = ?", id)
	if err != nil {
		return err
	}
	return nil
}
