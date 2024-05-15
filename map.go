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

// sensorData keeps track of the data for an uploaded piece of sensor data.
type sensorData struct {
	ID        uint64    `json:"id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Type      string    `json:"sensorType"`
	CSV       string    `json:"data"`
	Image     string    `json:"image"` // Not really hooked up. Probably should use something better than string.
	Date      time.Time `json:"timestamp"`
}

// mapHandler keeps track of sensor data displayed on the map.
// The handlers run in separate goroutines when getting network requests
// so care should be taken to avoid race conditions.
type mapHandler struct {
	id atomic.Uint64

	mu   sync.RWMutex
	data []*sensorData
}

// addSensorData is a handler that runs when the user uploads sensor data to the server.
func (s *mapHandler) addSensorData(c *gin.Context) {
	id := s.id.Add(1) - 1 // Increment directly but decrement locally so we get zero-indexing (and no race between getting the index and incrementing later).
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

// getMapCoordinates is a handler that runs when the website fetches map markers to draw.
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

// getMarker is a handler that runs when the website wants to get data for a specific marker.
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

// downloadHandler is a handler that runs when downloading sensor data.
func (s *mapHandler) downloadHandler(c *gin.Context) {
	markerID := c.Param("markerID")
	slog.Info("Initiated download of JSON data", slog.String("markerID", markerID))

	// Fake json data for the download:
	jsonData := []byte(`{
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
	}`)

	c.Data(http.StatusOK, "application/json", jsonData)
}

// bookmarkHandler is a handler that runs when the website checks if sensor data is bookmarked.
func (s *mapHandler) bookmarkHandler(c *gin.Context) {
	markerID := c.Param("markerID")
	userID := c.Param("userID")

	slog.Info("Checking if sensor data is bookmarked", slog.String("markerID", markerID), slog.String("userID", userID))

	// Fake data for bookmarking:
	c.JSON(http.StatusOK, gin.H{"bookmarked": markerID != "1"})
}
