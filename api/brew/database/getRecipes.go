package database

import (
	"pier/api/brew/types"
	"pier/storage"
)

func GetRecipes() ([]types.Recipe, error) {
	var recipes []types.Recipe

	db := storage.DB()

	rows, err := db.Query("SELECT * FROM `brew` ORDER BY `name` ASC")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var recipe types.Recipe
		err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.Coffee, &recipe.Water, &recipe.Grind, &recipe.Time, &recipe.Notes)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}
