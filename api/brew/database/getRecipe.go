package database

import (
	"pier/api/brew/database/model"
	"pier/storage"
)

func GetRecipe(id int64) (model.Recipe, error) {
	var recipe model.Recipe

	db, err := storage.DB()
	if err != nil {
		return recipe, err
	}

	res := db.First(&recipe, id)
	if res.Error != nil {
		return recipe, res.Error
	}

	return recipe, nil
}
