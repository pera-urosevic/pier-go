package reader

import (
	"net/http"
	"pier/api/reader/database"
	"pier/api/reader/types"
	"pier/lib"
	"strings"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) *gin.Engine {

	r.GET("/reader/bundles", func(c *gin.Context) {
		feeds, err := database.GetFeeds()
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		articles, err := database.GetArticles()
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		bundle := types.Bundle{Feeds: feeds, Articles: articles}
		c.JSON(http.StatusOK, bundle)
	})

	r.POST("/reader/feed/:name", func(c *gin.Context) {
		name := c.Param("name")
		feed, err := database.CreateFeed(name)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, feed)
	})

	r.GET("/reader/feed/:name", func(c *gin.Context) {
		name := c.Param("name")
		feed, err := database.GetFeed(name)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, feed)
	})

	r.PUT("/reader/feed/:name", func(c *gin.Context) {
		name := c.Param("name")
		var feed types.Feed
		err := c.BindJSON(&feed)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}
		if strings.HasPrefix(feed.Icon, "http") {
			bytes, err := lib.Download(feed.Icon)
			if err != nil {
				lib.GinError(c, err, true)
				return
			}
			feed.Icon = lib.Base64Encode(bytes)
		}
		feed, err = database.UpdateFeed(name, feed)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, feed)
	})

	r.DELETE("/reader/feed/:name", func(c *gin.Context) {
		name := c.Param("name")
		err := database.RemoveFeed(name)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, nil)
	})

	r.PUT("/reader/discard/feed/:name", func(c *gin.Context) {
		name := c.Param("name")
		err := database.DiscardFeed(name)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, nil)
	})

	r.PUT("/reader/discard/article/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := database.DiscardArticle(id)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, nil)
	})

	return r
}
