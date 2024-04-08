package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type SignUpForm struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
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

	errors := make(map[string]string)
	if form.FirstName == "" {
		errors["firstNameError"] = "First name is required"
	}
	if form.LastName == "" {
		errors["lastNameError"] = "Last name is required"
	}
	if form.Email == "" {
		errors["emailError"] = "Email is required"
	}
	if form.Password == "" {
		errors["passwordError"] = "Password is required"
	}

	if len(errors) > 0 {
		errors["errorMsg"] = "Your registration failed! Please check the form."
		c.JSON(http.StatusBadRequest, errors)
		return
	}

	fmt.Println("Super secure form")
	fmt.Println("Firstname: ", form.FirstName)
	fmt.Println("Lastname: ", form.LastName)
	fmt.Println("Email: ", form.Email)
	fmt.Println("Password: ", form.Password)

	// Include a custom message in the response
	c.JSON(http.StatusOK, gin.H{
		"status":     "Form submitted fine",
		"successMsg": "Your registration is successful!",
	})
}

func main() {
	fmt.Println("Hello from BLHub server!")

	engine := gin.Default()

	// Enable CORS because we cant run frontend and backend on same port
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // Frontend url:port
	engine.Use(cors.New(config))

	engine.GET("/index.html", helloHandler)
	engine.POST("/name", postHandler)
	engine.POST("/signup", signUpHandler)

	err := engine.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
