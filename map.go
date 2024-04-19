package main

import (
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
	CSV       string    `json:"data"`
	Image     string    `json:"image"` // Not really hooked up. Probably should use something betetr than string.
	Date      time.Time `json:"timestamp"`
}

type mapHandler struct {
	id atomic.Uint64

	mu   sync.RWMutex
	data []*sensorData
}

func (s *mapHandler) addSensorData(c *gin.Context) {
	id := s.id.Add(1) - 1 // Increment directly but decrement so we get zero-indexing.
	form := &sensorData{ID: id}

	if err := c.ShouldBindJSON(form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		slog.Error("Failed to parse form", slog.String("error", err.Error()))
		return
	}

	s.mu.Lock()
	s.data = append(s.data, form)
	s.mu.Unlock()

	c.JSON(http.StatusOK, "Sensor data was added successfully")

	slog.Info("Stored a new sensor data", "id", form.ID, "timestamp", form.Date)
}

func (s *mapHandler) getMapCoordinates(c *gin.Context) {
	type coordinateData struct {
		ID        uint64  `json:"id"`
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lng"`
	}

	s.mu.RLock()
	markers := make([]coordinateData, len(s.data))
	for i, data := range s.data {
		markers[i] = coordinateData{
			ID:        data.ID,
			Latitude:  data.Latitude,
			Longitude: data.Longitude,
		}
	}
	s.mu.RUnlock()

	/*
		// Fake coordinates data for the map:
		markers = []coordinateData{
			{ID: 1, Latitude: 30.1695, Longitude: 10.9354},
			{ID: 2, Latitude: 61.1695, Longitude: 24.9354},
			{ID: 3, Latitude: 63.1695, Longitude: 24.9354},
			{ID: 4, Latitude: 64.1695, Longitude: 24.9354},
			{ID: 5, Latitude: 65.1695, Longitude: 24.9354},
		}
	*/

	c.JSON(http.StatusOK, gin.H{"markers": markers})

	slog.Info("Sent over all map coordinates")
}

func (s *mapHandler) getMarker(c *gin.Context) {
	param := c.Param("markerID")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusNotFound, "Incorrect marker ID")
		slog.Error("Could not parse marker ID", slog.String("id", param))
		return
	}

	s.mu.RLock()
	data := s.data[id]
	s.mu.RUnlock()

	// Fake sensor types data with unit:
	sensorTypes := []map[string]string{
		{"type": "temperature", "unit": "°C"},
		{"type": "humidity", "unit": "%"},
		{"type": "light", "unit": "lux"},
		{"type": "moisture", "unit": "%"},
	}

	c.JSON(http.StatusOK, gin.H{
		"sensorID":    id,
		"sensorName":  data.Type + " sensor",
		"sensorImage": "https://www.plantagen.se/dw/image/v2/BCMR_PRD/on/demandware.static/-/Library-Sites-PlantagenShared/default/dw258d02d2/1000/elefantore-pilea-peperomioides.jpg?sw=1024",
		"sensorTypes": sensorTypes,
		"lastUpdated": data.Date.Format(time.DateTime),
	})

	slog.Info("Sent over data for a marker", "id", data.ID, "timestamp", data.Date)
}
