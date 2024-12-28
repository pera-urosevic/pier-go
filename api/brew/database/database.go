package database

import (
	"pier/api/brew/types"
	"pier/storage"
)

func GetRecipes() ([]types.Recipe, error) {
	db := storage.DB()
	rows, err := db.Query("SELECT * FROM `brew` ORDER BY `name` ASC")
	if err != nil {
		return nil, err
	}
	var recipes []types.Recipe
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

func GetRecipe(id int64) (types.Recipe, error) {
	db := storage.DB()
	row := db.QueryRow("SELECT * FROM `brew` WHERE `id` = ?", id)
	var recipe types.Recipe
	err := row.Scan(&recipe.ID, &recipe.Name, &recipe.Coffee, &recipe.Water, &recipe.Grind, &recipe.Time, &recipe.Notes)
	if err != nil {
		return types.Recipe{}, err
	}
	return recipe, nil
}

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

func RemoveRecipe(id int64) error {
	db := storage.DB()
	_, err := db.Exec("DELETE FROM `brew` WHERE `id` = ?", id)
	if err != nil {
		return err
	}
	return nil
}
