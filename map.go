package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
)

type sensorData struct {
	ID        uint64    `json:"id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Type      string    `json:"sensorType"`
	Data      string    `json:"data"`
	Image     string    `json:"image"`
	Date      time.Time `json:"timestamp"`
}

type mapHandler struct {
	id atomic.Uint64

	mu   sync.RWMutex
	data []*sensorData
}

func (s *mapHandler) addSensorData(c *gin.Context) {
	form := &sensorData{ID: s.id.Add(1)}

	if err := c.ShouldBindJSON(form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		slog.Error("Failed to parse form", slog.String("error", err.Error()))
		return
	}

	fmt.Println(form)

	s.mu.Lock()
	s.data = append(s.data, form)
	s.mu.Unlock()
}
func (s *mapHandler) getMapCoordinates(c *gin.Context) {
	fmt.Println("FETCH COORDINATES")

	s.mu.RLock()
	c.JSON(http.StatusOK, s.data)
	s.mu.RUnlock()
}

func (s *mapHandler) getMarker(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("markerID"))
	fmt.Println("Marker ID: ", id)

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
		"sensorID":    id,
		"sensorName":  "SENSOR EXAMPLE NAME",
		"sensorImage": "https://www.plantagen.se/dw/image/v2/BCMR_PRD/on/demandware.static/-/Library-Sites-PlantagenShared/default/dw258d02d2/1000/elefantore-pilea-peperomioides.jpg?sw=1024",
		"sensorTypes": sensorTypes,
		"lastUpdated": currentTime,
	})

	/*
		s.mu.RLock()
		data := s.data[id]
		s.mu.RUnlock()

		c.JSON(http.StatusOK, data)
	*/
}
