package tvshows

import (
	"net/http"
	"pier/api/tvshows/database"
	"pier/api/tvshows/tvmaze"
	"pier/api/tvshows/types"
	"pier/lib"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) *gin.Engine {

	r.GET("/tvshows", func(c *gin.Context) {
		tvshows, err := database.GetTVShows()
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, tvshows)
	})

	r.GET("/tvshows/tvshow/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}
		tvshows, err := database.GetTVShow(id)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, tvshows)
	})

	r.POST("/tvshows/tvshow", func(c *gin.Context) {
		var tvshow types.TVShow
		err := c.BindJSON(&tvshow)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}
		if tvshow.Image != nil {
			if strings.HasPrefix(*tvshow.Image, "http") {
				bytes, err := lib.Download(*tvshow.Image)
				if err != nil {
					lib.GinError(c, err, true)
					return
				}
				base64 := lib.Base64Encode(bytes)
				tvshow.Image = &base64
			}
		}
		tvshow, err = database.CreateTVShow(tvshow)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, tvshow)
	})

	r.PUT("/tvshows/tvshow/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}
		var tvshow types.TVShow
		err = c.BindJSON(&tvshow)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}
		err = database.UpdateTVShow(id, tvshow)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, tvshow)
	})

	r.DELETE("/tvshows/tvshow/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}
		err = database.RemoveTVShow(id)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, nil)
	})

	r.GET("/tvshows/tvmaze/search/:title", func(c *gin.Context) {
		title := c.Param("title")
		results, err := tvmaze.Search(title)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, results)
	})

	r.POST("/tvshows/tvmaze/get/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			lib.GinError(c, err, false)
		}
		result, err := tvmaze.Get(id)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, result)
	})

	return r
}
