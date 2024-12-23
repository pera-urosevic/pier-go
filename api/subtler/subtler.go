package subtler

import (
	"net/http"
	"os"
	"path/filepath"
	"pier/api/subtler/extract"
	"pier/api/subtler/types"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) *gin.Engine {

	r.GET("/subtler/entries/*path", func(c *gin.Context) {
		path := os.Getenv("SUBTLER_ROOT")
		path = filepath.Join(path, c.Param("path"))
		dir, err := os.ReadDir(path)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		entries := []types.Entry{}
		for _, f := range dir {
			entry := types.Entry{Name: f.Name(), Dir: f.IsDir()}
			entries = append(entries, entry)
		}
		c.JSON(http.StatusOK, entries)
	})

	r.GET("/subtler/extract/*video", func(c *gin.Context) {
		video := os.Getenv("SUBTLER_ROOT")
		video = filepath.Join(video, c.Param("video"))
		log, err := extract.Extract(video)
		if err != nil {
			payload := map[string]string{"error": err.Error(), "log": log}
			c.JSON(http.StatusInternalServerError, payload)
			return
		}
		c.JSON(http.StatusOK, log)
	})

	return r
}
