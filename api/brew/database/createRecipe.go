package database

import (
	"pier/api/brew/types"
	"pier/storage"
)

func CreateRecipe(recipe types.Recipe) (types.Recipe, error) {
	db := storage.DB()

	res, err := db.Exec("INSERT INTO `brew` (`name`, `coffee`, `water`, `grind`, `time`, `notes`) VALUES (?, ?, ?, ?, ?, ?)", recipe.Name, recipe.Coffee, recipe.Water, recipe.Grind, recipe.Time, recipe.Notes)
	if err != nil {
		return recipe, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return recipe, err
	}

	recipe.ID = id
	return recipe, nil
}
