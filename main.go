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
	sloggin "github.com/samber/slog-gin"
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

	// Log with default gin logger and a middleware for log/slog.
	// This makes sure that we get the fancy stdout logs but also log to the log file.
	defaultGinLogger := gin.Logger()
	slogginLogger := sloggin.New(slog.Default())
	ginLogger := func(ctx *gin.Context) {
		defaultGinLogger(ctx)
		slogginLogger(ctx)
	}
	engine.Use(ginLogger, gin.Recovery())

	// Enable CORS because we cant run frontend and backend on same port
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // Frontend url:port
	engine.Use(cors.New(config))

	engine.POST("/signup", signUpHandler)

	err := engine.Run(":8080")
	if err != nil {
		slog.Error("Gin router encountered an error", slog.String("error", err.Error()))
	}
}
