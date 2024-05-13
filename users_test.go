package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsersSignup(t *testing.T) {
	router := setupRounter()

	// Test creating an initial user:
	user := &user{
		FirstName: "Anders",
		LastName:  "Andersson",
		Email:     "anders@anders.com",
		Password:  "secret",
	}

	body, _ := json.Marshal(user)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(body))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code) // Should work fine.

	// Test creating a new user with the same email but different name:
	user.FirstName = "Bosse"
	body, _ = json.Marshal(user)

	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(body))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotAcceptable, w.Code) // Should error about user not being accepted.

	// Test passing invalid data:
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader("broken"))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code) // Should error about failing to parse form.

	// Test using incorrect request type:
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/signup", strings.NewReader("broken"))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code) // Should error with 404 because Get is not handled.
}

func TestUsersLogin(t *testing.T) {
	router := setupRounter()

	// Test creating an initial user:
	user := &user{
		FirstName: "Anders",
		LastName:  "Andersson",
		Email:     "anders@anders.com",
		Password:  "secret",
	}

	body, _ := json.Marshal(user)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(body))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code) // Should work fine.

	// Test logging in with the created user:
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code) // Should work fine.

	// Test logging in with the created user but wrong password:
	user.Password = "wrong"
	body, _ = json.Marshal(user)
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotAcceptable, w.Code) // Should error about wrong password.

	// Test logging in with a not created user:
	user.Email = "not@registered.se"
	body, _ = json.Marshal(user)
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotAcceptable, w.Code) // Should error about not being registered.

	// Test passing invalid data:
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/login", strings.NewReader("broken"))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code) // Should error about failing to parse form.
}
