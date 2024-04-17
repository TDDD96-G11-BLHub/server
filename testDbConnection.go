package main

import (
	"fmt"
	"net/http"

	"github.com/TDDD96-G11-BLHub/dbman/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func testDbConnection(client *mongo.Client) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		fmt.Print("testDbConnection")
		db.TestConnection(client)

		c.JSON(http.StatusOK, gin.H{})
	}

	return gin.HandlerFunc(fn)
}
