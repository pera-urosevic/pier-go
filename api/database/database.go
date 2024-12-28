package database

import (
	"net/http"
	"pier/api/database/database"
	"pier/api/database/types"
	"pier/lib"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) *gin.Engine {

	r.GET("/database/index", func(c *gin.Context) {
		databaseIndex, err := database.GetDatabaseIndex()
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, databaseIndex)
	})

	r.GET("/database/:database/collection/:collection", func(c *gin.Context) {
		databaseName := c.Param("database")
		collectionName := c.Param("collection")
		where := c.Query("where")
		if where == "" {
			where = "true"
		}
		order := c.Query("order")
		if order == "" {
			order = "id DESC"
		}
		offset, err := strconv.Atoi(c.Query("offset"))
		if err != nil {
			offset = 0
		}
		docs, err := database.GetDocs(databaseName, collectionName, where, order, offset)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, docs)
	})

	r.GET("/database/:database/collection/:collection/facets", func(c *gin.Context) {
		databaseName := c.Param("database")
		collectionName := c.Param("collection")
		facets, err := database.GetFacets(databaseName, collectionName)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, facets)
	})

	r.PUT("/database/:database/collection/:collection/facets", func(c *gin.Context) {
		databaseName := c.Param("database")
		collectionName := c.Param("collection")
		var facets string
		err := c.BindJSON(&facets)
		if err != nil {
			lib.GinError(c, err, false)
		}
		err = database.UpdateFacets(databaseName, collectionName, facets)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, facets)
	})

	r.GET("/database/:database/collection/:collection/count", func(c *gin.Context) {
		databaseName := c.Param("database")
		collectionName := c.Param("collection")
		where := c.Query("where")
		where = strings.ReplaceAll(where, "|", "%")
		if where == "" {
			where = "true"
		}
		count, err := database.GetDocsCount(databaseName, collectionName, where)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, count)
	})

	r.POST("/database/:database/collection/:collection/autocompletes", func(c *gin.Context) {
		databaseName := c.Param("database")
		collectionName := c.Param("collection")
		var autocompleteFields types.AutocompleteFields
		err := c.BindJSON(&autocompleteFields)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}
		autocompletes, err := database.GetAutocompletes(databaseName, collectionName, autocompleteFields)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, autocompletes)
	})

	r.GET("/database/:database/collection/:collection/doc/:id", func(c *gin.Context) {
		databaseName := c.Param("database")
		collectionName := c.Param("collection")
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}
		doc, err := database.GetDoc(databaseName, collectionName, id)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, doc)
	})

	r.GET("/database/:database/collection/:collection/duplicate/:id/:name", func(c *gin.Context) {
		databaseName := c.Param("database")
		collectionName := c.Param("collection")
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}
		name := c.Param("name")
		exists, err := database.GetDuplicate(databaseName, collectionName, id, name)
		if err != nil {
			lib.GinError(c, err, true)
		}
		c.JSON(http.StatusOK, exists)
	})

	r.POST("/database/:database/collection/:collection/doc", func(c *gin.Context) {
		databaseName := c.Param("database")
		collectionName := c.Param("collection")
		var doc types.Doc
		err := c.BindJSON(&doc)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}
		id, err := database.CreateDoc(databaseName, collectionName, doc)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, id)
	})

	r.PUT("/database/:database/collection/:collection/doc/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}
		databaseName := c.Param("database")
		collectionName := c.Param("collection")
		var doc types.Doc
		err = c.BindJSON(&doc)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}
		err = database.UpdateDoc(databaseName, collectionName, id, doc)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		docNew, err := database.GetDoc(databaseName, collectionName, id)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, docNew)
	})

	r.DELETE("/database/:database/collection/:collection/doc/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}
		databaseName := c.Param("database")
		collectionName := c.Param("collection")
		err = database.RemoveDoc(databaseName, collectionName, id)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, nil)
	})

	r.POST("/database/cache/:key", func(c *gin.Context) {
		key := c.Param("key")
		var value string
		err := c.BindJSON(&value)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}
		err = database.SetCache(key, value)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
	})

	r.GET("/database/cache/:key", func(c *gin.Context) {
		key := c.Param("key")
		value, err := database.GetCache(key)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, value)
	})

	r.DELETE("/database/cache/:key", func(c *gin.Context) {
		key := c.Param("key")
		err := database.RemoveCache(key)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, nil)
	})

	return r
}
