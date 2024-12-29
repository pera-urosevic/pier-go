package database

import (
	"pier/api/brew/types"
	"pier/storage"
)

func GetRecipe(id int64) (types.Recipe, error) {
	var recipe types.Recipe

	db := storage.DB()

	row := db.QueryRow("SELECT * FROM `brew` WHERE `id` = ?", id)
	err := row.Scan(&recipe.ID, &recipe.Name, &recipe.Coffee, &recipe.Water, &recipe.Grind, &recipe.Time, &recipe.Notes)
	if err != nil {
		return types.Recipe{}, err
	}

	return recipe, nil
}
