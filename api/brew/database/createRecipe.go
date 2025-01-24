package database

import (
	"pier/api/brew/database/model"
	"pier/storage"
)

func CreateRecipe(recipe model.Recipe) (model.Recipe, error) {
	db, err := storage.DB()
	if err != nil {
		return recipe, err
	}

	res := db.Create(&recipe)
	if res.Error != nil {
		return recipe, res.Error
	}

	return recipe, nil
}
