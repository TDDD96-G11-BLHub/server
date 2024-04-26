package main

import (
	"fmt"
	"net/http"

	"github.com/TDDD96-G11-BLHub/dbman/db"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
)

func mapHandler(client *mongo.Client) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		fmt.Println("FETCH COORDINATES ")

		db.TestConnection(client)

		// Fake coordinates data for the map
		markers := []map[string]float64{
			{"id": 1, "lat": 30.1695, "lng": 10.9354},
			{"id": 2, "lat": 61.1695, "lng": 24.9354},
			{"id": 3, "lat": 63.1695, "lng": 24.9354},
			{"id": 4, "lat": 64.1695, "lng": 24.9354},
			{"id": 5, "lat": 65.1695, "lng": 24.9354},
		}

		// Include a custom message in the response
		c.JSON(http.StatusOK, gin.H{
			"markers": markers,
		})
	}

	return gin.HandlerFunc(fn)
}

func markerHandler(c *gin.Context) {
	markerID := c.Param("markerID")
	fmt.Println("Marker ID: ", markerID)

	// Include a custom message in the response
	c.JSON(http.StatusOK, gin.H{
		"sensorName":  "TEST NAME",
		"image":       "https://www.plantagen.se/dw/image/v2/BCMR_PRD/on/demandware.static/-/Library-Sites-PlantagenShared/default/dw258d02d2/1000/elefantore-pilea-peperomioides.jpg?sw=1024",
		"sensorTypes": "TEST TYPES",
		"lastUpdate":  "TEST UPDATE",
	})
}
