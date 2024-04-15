package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"slices"
	"sync"

	"github.com/gin-gonic/gin"
)

type user struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// LogValue formats the user as a logged value and hides the password field.
func (u *user) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("email", u.Email),
		slog.String("firstname", u.FirstName),
		slog.String("lastname", u.LastName),
	)
}

type userHandler struct {
	users []*user
	mu    sync.RWMutex
}

func (u *userHandler) login(c *gin.Context) {
	type loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	form := &loginData{}

	if err := c.ShouldBindJSON(form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		slog.Error("Failed to parse login form", slog.String("error", err.Error()))
		return
	}

	fmt.Println(form)

	u.mu.RLock()
	contains := slices.ContainsFunc(u.users, func(item *user) bool { return item.Email == form.Email && item.Password == form.Password })
	u.mu.RUnlock()

	if !contains {
		c.JSON(http.StatusNotAcceptable, "A user with the specified email address already exists!")
		slog.Warn("A user tried to login with the wrong email", slog.String("email", form.Email))
		return
	}

	slog.Info("A new user logged in", slog.String("email", form.Email))

	c.JSON(http.StatusOK, gin.H{
		"status":     "Form submitted fine",
		"successMsg": "Your login is successful!",
	})
}

func (u *userHandler) signup(c *gin.Context) {
	form := &user{}

	if err := c.ShouldBindJSON(form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		slog.Error("Failed to parse signup form", slog.String("error", err.Error()))
		return
	}

	u.mu.RLock()
	contains := slices.ContainsFunc(u.users, func(item *user) bool { return item.Email == form.Email })
	u.mu.RUnlock()

	if contains {
		c.JSON(http.StatusNotAcceptable, "A user with the specified email address already exists!")
		slog.Warn("A user tried to use the same email twice", "user", form)
		return
	}

	u.mu.Lock()
	u.users = append(u.users, form)
	u.mu.Unlock()

	slog.Info("A new user was created", "user", form)

	// Include a custom message in the response
	c.JSON(http.StatusOK, gin.H{
		"status":     "Form submitted fine",
		"successMsg": "Your registration is successful!",
	})
}
