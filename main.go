package main

//Test

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type SignUpForm struct {
	FirstName    string `json:"firstname" binding:"required"`
	LastName    string `json:"lastname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func helloHandler(c *gin.Context) {
	fmt.Fprintln(c.Writer,
		`<!DOCTYPE html>
	<html>
	<form method="post" action="/name">
    <label for="firstname">First name:</label>
    <input type="text" name="firstname" /><br />
    <label for="lastname">Last name:</label>
    <input type="text" name="lastname" /><br />
    <input type="submit" />
	</form></html>`)
}

func postHandler(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	name := c.Request.Form.Get("firstname")
	lastname := c.Request.Form.Get("lastname")
	fmt.Println(name, lastname)

	fmt.Fprintln(c.Writer, "Form submitted fine")
}

func signUpHandler(c *gin.Context) {
	var form SignUpForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("ERROR ")
		return
	}

	fmt.Println("Super secure form ;)")
	fmt.Println("Firstname: ", form.FirstName)
	fmt.Println("Lastname: ", form.LastName)
	fmt.Println("Email: ", form.Email)
	fmt.Println("Password: ", form.Password)

	c.JSON(http.StatusOK, gin.H{"status": "Form submitted fine"})
}

func main() {
	fmt.Println("Hello from BLHub server!")

	engine := gin.Default()

	// Enable CORS because we cant run frontend and backend on same port
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"} // Frontend url:port
	engine.Use(cors.New(config))

	engine.GET("/index.html", helloHandler)
	engine.POST("/name", postHandler)
	engine.POST("/signup", signUpHandler)

	err := engine.Run(":5000")
	if err != nil {
		log.Fatal(err)
	}
}
