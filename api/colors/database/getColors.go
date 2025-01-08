package database

import (
	"pier/api/colors/types"
	"pier/storage"
)

func GetColors() ([]types.Color, error) {
	var colors = []types.Color{}

	db := storage.DB()

	rows, err := db.Query("SELECT * FROM `colors` ORDER BY `name` ASC")
	if err != nil {
		return colors, err
	}

	for rows.Next() {
		var color types.Color
		err := rows.Scan(&color.Name, &color.H, &color.S, &color.L)
		if err != nil {
			return colors, err
		}
		colors = append(colors, color)
	}

	return colors, nil
}
