package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func bookmarkHandler(c *gin.Context) {
	fmt.Println("BOOKMARK MARKER:")

	markerID := c.Param("markerID")
	userID := c.Param("userID")
	fmt.Println("Marker ID: ", markerID)
	fmt.Println("User ID: ", userID)

	if markerID == "1" {
		c.JSON(http.StatusOK, gin.H{
			"bookmarked": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"bookmarked": true,
	})
}
