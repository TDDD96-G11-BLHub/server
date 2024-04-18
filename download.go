package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func downloadHandler(c *gin.Context) {
	fmt.Println("DOWNLOAD JSON FILE")

	markerID := c.Param("markerID")
	fmt.Println("Marker ID: ", markerID)

	// Fake json data for the download
	jsonData := `{
		"sensorID": 1,
		"sensorName": "SENSOR EXAMPLE NAME",
		"sensorImage": "https://www.plantagen.se/dw/image/v2/BCMR_PRD/on/demandware.static/-/Library-Sites-PlantagenShared/default/dw258d02d2/1000/elefantore-pilea-peperomioides.jpg?sw=1024",
		"sensorTypes": [
			{"type": "temperature", "unit": "°C"},
			{"type": "humidity", "unit": "%"},
			{"type": "light", "unit": "lux"},
			{"type": "moisture", "unit": "%"}
		],
		"lastUpdated": "2021-01-01 12:00:00"
	}`

	c.Data(http.StatusOK, "application/json", []byte(jsonData))
}
