package main

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	err := os.Mkdir("logs", 0o777)
	if err != nil && !errors.Is(err, os.ErrExist) {
		slog.Warn("Failed to create log", err)
		return
	}

	filename := time.Now().Format("logs/" + time.DateTime + ".txt")
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o666)
	if err != nil {
		slog.Warn("Failed to open log file", err)
		return
	}

	writers := io.MultiWriter(file, os.Stderr)
	opts := &slog.HandlerOptions{AddSource: true}
	logger := slog.New(slog.NewTextHandler(writers, opts))
	slog.SetDefault(logger)
}

func main() {
	fmt.Println("Starting up BLHub server!")

	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	// Enable CORS because we cant run frontend and backend on same port
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // Frontend url:port
	engine.Use(cors.New(config))

	users := &userHandler{}

	// User handlers:
	engine.POST("/signup", users.signup)
	engine.POST("/login", users.login)

	mapData := &mapHandler{}

	// Handlers related to map data:
	engine.POST("/sensordata/add", mapData.addSensorData)
	engine.GET("/map", mapData.getMapCoordinates)
	engine.GET("/map/:markerID", mapData.getMarker)
	engine.GET("/download/:markerID", mapData.downloadHandler)
	engine.GET("/bookmark/:markerID/:userID", mapData.bookmarkHandler)

	err := engine.Run(":8080")
	if err != nil {
		slog.Error("Gin router encountered an error", slog.String("error", err.Error()))
	}
}
