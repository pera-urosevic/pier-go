package monitor

import (
	"net/http"
	"pier/api/monitor/database"
	"pier/lib"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) *gin.Engine {

	r.GET("/monitor", func(c *gin.Context) {
		monitorData, err := database.GetData()
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, monitorData)
	})

	r.DELETE("/monitor/notification/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}
		err = database.RemoveNotification(id)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, nil)
	})

	r.DELETE("/monitor/notifications/:channel", func(c *gin.Context) {
		channel := c.Param("channel")
		err := database.RemoveNotifications(channel)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, nil)
	})

	return r
}
