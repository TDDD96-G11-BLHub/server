package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting up BLHub server!")

	engine := gin.Default()

	// Enable CORS because we cant run frontend and backend on same port
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // Frontend url:port
	engine.Use(cors.New(config))

	engine.POST("/signup", signUpHandler)

	err := engine.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
