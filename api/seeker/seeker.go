package seeker

import (
	"net/http"
	"pier/api/seeker/database"
	"pier/api/seeker/types"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) *gin.Engine {

	r.GET("/seeker/targets", func(c *gin.Context) {
		targets, err := database.GetTargets()
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, targets)
	})

	r.POST("/seeker/target", func(c *gin.Context) {
		var target types.Target
		err := c.BindJSON(&target)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}
		err = database.CreateTarget(target)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		targets, err := database.GetTargets()
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, targets)
	})

	r.PUT("/seeker/target/:title", func(c *gin.Context) {
		title := c.Param("title")
		var target types.Target
		err := c.BindJSON(&target)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}
		err = database.UpdateTarget(title, target)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		targets, err := database.GetTargets()
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, targets)
	})

	r.DELETE("/seeker/target/:title", func(c *gin.Context) {
		title := c.Param("title")
		err := database.RemoveTarget(title)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		targets, err := database.GetTargets()
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, targets)
	})

	return r
}
