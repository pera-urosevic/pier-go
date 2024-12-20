package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinError(c *gin.Context, err error, internal bool) {
	if internal {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
