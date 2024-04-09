package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignUpForm struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
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
