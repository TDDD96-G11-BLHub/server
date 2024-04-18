package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func mapHandler(c *gin.Context) {
	fmt.Println("FETCH COORDINATES ")

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

	// fake sensor types data with unit
	sensorTypes := []map[string]string{
		{"type": "temperature", "unit": "°C"},
		{"type": "humidity", "unit": "%"},
		{"type": "light", "unit": "lux"},
		{"type": "moisture", "unit": "%"},
	}

	// current time
	currentTime := time.Now().Format(time.DateTime)

	// Include a custom message in the response
	c.JSON(http.StatusOK, gin.H{
		"sensorID":    markerID,
		"sensorName":  "SENSOR EXAMPLE NAME",
		"sensorImage": "https://www.plantagen.se/dw/image/v2/BCMR_PRD/on/demandware.static/-/Library-Sites-PlantagenShared/default/dw258d02d2/1000/elefantore-pilea-peperomioides.jpg?sw=1024",
		"sensorTypes": sensorTypes,
		"lastUpdated": currentTime,
	})
}
