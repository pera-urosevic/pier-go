package brew

import (
	"net/http"
	"pier/api/brew/database"
	"pier/api/brew/types"
	"pier/lib"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) *gin.Engine {

	r.GET("/brew/recipes", func(c *gin.Context) {
		recipes, err := database.GetRecipes()
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, recipes)
	})

	r.POST("/brew/recipe", func(c *gin.Context) {
		var recipe types.Recipe
		err := c.BindJSON(&recipe)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}
		recipe, err = database.CreateRecipe(recipe)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, recipe)
	})

	r.PUT("/brew/recipe", func(c *gin.Context) {
		var recipe types.Recipe
		err := c.BindJSON(&recipe)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}
		recipe, err = database.UpdateRecipe(recipe)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, recipe)
	})

	r.GET("/brew/recipe/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}
		recipe, err := database.GetRecipe(id)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, recipe)
	})

	r.DELETE("/brew/recipe/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}
		err = database.RemoveRecipe(id)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, nil)
	})

	return r
}
