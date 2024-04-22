package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func graphHandlerTest(c *gin.Context) {
	fmt.Println("FETCH DATAPOINTS ")

	sensordata := []map[string]interface{}{
		{
			"Time":  "15:40:21",
			"Roll":  0.438369,
			"Pitch": -1.005255,
		},
		{
			"Time":  "15:40:21",
			"Roll":  0.789288,
			"Pitch": -1.997637,
		},
		{
			"Time":  "15:40:22",
			"Roll":  0.945767,
			"Pitch": -2.935617,
		},
		{
			"Time":  "15:40:22",
			"Roll":  1.032266,
			"Pitch": -3.588959,
		},
		{
			"Time":  "15:40:22",
			"Roll":  0.874982,
			"Pitch": -3.853505,
		},
	}

	// Include a custom message in the response
	c.JSON(http.StatusOK, gin.H{
		"sensordata": sensordata,
	})
}

/*
	Function to be used with marker ID

func graphHandler(c *gin.Context) {
	markerID := c.Param("markerID")
	fmt.Println("Marker ID: ", markerID)

	dataPoints := []map[string]interface{}{
		{
			"Time":  "15:40:21",
			"Roll":  0.438369,
			"Pitch": -1.005255,
			"Yaw":   0.31074,
		},
		{
			"Time":  "15:40:21",
			"Roll":  0.789288,
			"Pitch": -1.997637,
			"Yaw":   0.745679,
		},
		{
			"Time":  "15:40:22",
			"Roll":  0.945767,
			"Pitch": -2.935617,
			"Yaw":   1.368871,
		},
		{
			"Time":  "15:40:22",
			"Roll":  1.032266,
			"Pitch": -3.588959,
			"Yaw":   2.295047,
		},
		{
			"Time":  "15:40:22",
			"Roll":  0.874982,
			"Pitch": -3.853505,
			"Yaw":   3.402261,
		},
	}
	fmt.Println(dataPoints)
	// Include a custom message in the response
	c.JSON(http.StatusOK, gin.H{
		"sensorName":  "TEST NAME",
		"image":       "https://www.plantagen.se/dw/image/v2/BCMR_PRD/on/demandware.static/-/Library-Sites-PlantagenShared/default/dw258d02d2/1000/elefantore-pilea-peperomioides.jpg?sw=1024",
		"sensorTypes": "TEST TYPES",
		"lastUpdate":  "TEST UPDATE",
	})
}
*/
