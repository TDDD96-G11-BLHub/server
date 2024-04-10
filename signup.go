package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	feedbackNoFirstName = "First name is required"
	feedbackNoLastName  = "Last name is required"
	feedbackNoEmail     = "Email is requried"
	feedbackNoPassword  = "Password is requried"
	feedbackFormMessage = "Your registration failed! Please check the form."
)

type signupFormData struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type signupFormFeedback struct {
	FirstName string `json:"firstNameError"`
	LastName  string `json:"lastNameError"`
	Email     string `json:"emailError"`
	Password  string `json:"passwordError"`
	Message   string `json:"errorMsg"`
}

func signUpHandler(c *gin.Context) {
	form := signupFormData{}

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		slog.Error("Failed to parse signup form", slog.String("error", err.Error()))
		return
	}

	errors := signupFormFeedback{}
	if form.FirstName == "" {
		errors.FirstName = feedbackNoFirstName
	}
	if form.LastName == "" {
		errors.LastName = feedbackNoLastName
	}
	if form.Email == "" {
		errors.Email = feedbackNoEmail
	}
	if form.Password == "" {
		errors.Password = feedbackNoPassword
	}

	emptyErrors := signupFormFeedback{}
	if errors != emptyErrors {
		errors.Message = feedbackFormMessage
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
