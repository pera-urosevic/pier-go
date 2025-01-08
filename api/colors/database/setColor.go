package database

import (
	"pier/api/colors/types"
	"pier/storage"
)

func SetColor(color types.Color) error {
	db := storage.DB()

	_, err := db.Exec("INSERT INTO `colors` (`name`, `h`, `s`, `l`) VALUES (?, ?, ?, ?)", color.Name, color.H, color.S, color.L)
	if err != nil {
		return err
	}

	return nil
}
