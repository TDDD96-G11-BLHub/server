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

func loginHandler(c *gin.Context) {
	var form loginForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("ERROR ")
		return
	}

	var testEmail string = "test@test.test"
	var testPassword string = "testtest"

	if testEmail != form.Email || testPassword != form.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"errorMsg": "Your login failed! Something something on the dark side",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"successMsg": "Your login is successful!",
	})

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
	engine.POST("/signup", users.signup)
	engine.POST("/login", users.login)

	err := engine.Run(":8080")
	if err != nil {
		slog.Error("Gin router encountered an error", slog.String("error", err.Error()))
	}
}
