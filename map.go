package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func mapHandler(c *gin.Context) {
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

func main() {
	fmt.Println("Hello from BLHub server!")

	engine := gin.Default()

	// Enable CORS because we cant run frontend and backend on same port
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // Frontend url:port
	engine.Use(cors.New(config))

	engine.GET("/map", mapHandler)
	engine.GET("/map/:markerID", markerHandler)

	err := engine.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
