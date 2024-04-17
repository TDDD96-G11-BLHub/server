package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	// Create connection to database
	if errdb := godotenv.Load(); errdb != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, errdb := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if errdb != nil {
		panic(errdb)
	}
	defer func() {
		if errdb := client.Disconnect(context.TODO()); errdb != nil {
			panic(errdb)
		}
	}()

	users := &userHandler{}
	engine.POST("/signup", users.signup)
	engine.POST("/login", users.login)

	// Map handlers
	engine.GET("/map", mapHandler)
	engine.GET("/map/:markerID", markerHandler)

	engine.GET("/testDbConnection", testDbConnection(client))

	err := engine.Run(":8080")
	if err != nil {
		slog.Error("Gin router encountered an error", slog.String("error", err.Error()))
	}
}
