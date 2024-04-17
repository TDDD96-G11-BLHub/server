package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func addSensorHandler(c *gin.Context) {
	type sensorData struct {
		Longitude float64   `json:"longitude"`
		Latitude  float64   `json:"latitude"`
		Type      int       `json:"type"`
		Data      string    `json:"file"`
		Image     string    `json:"image"`
		Date      time.Time `json:"timestamp"`
	}

	form := &sensorData{}

	if err := c.ShouldBindJSON(form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		slog.Error("Failed to parse form", slog.String("error", err.Error()))
		return
	}

	fmt.Println(form)
}
