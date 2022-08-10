package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPixelHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
	// to do add get pixels func
}
