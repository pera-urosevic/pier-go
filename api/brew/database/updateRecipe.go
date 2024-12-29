package database

import (
	"pier/api/brew/types"
	"pier/storage"
)

func UpdateRecipe(recipe types.Recipe) (types.Recipe, error) {
	db := storage.DB()

	res, err := db.Exec("UPDATE `brew` SET `name` = ?, `coffee` = ?, `water` = ?, `grind` = ?, `time` = ?, `notes` = ? WHERE `id` = ?", recipe.Name, recipe.Coffee, recipe.Water, recipe.Grind, recipe.Time, recipe.Notes, recipe.ID)
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
