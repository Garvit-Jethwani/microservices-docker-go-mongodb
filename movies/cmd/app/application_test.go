// ********RoostGPT********
/*
Application Test generated by RoostGPT for test ApplicationTest-Golang-2 using AI Type Open AI and AI Model gpt-4-turbo


*/

// ********RoostGPT********
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestAllMovies(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/movies/", nil)
	assert.NoError(t, err)
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(app.all)
	handler.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestFindByID(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies/{id}", app.findByID)
	req, err := http.NewRequest("GET", "/api/movies/123", nil)
	assert.NoError(t, err)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	expectedStatus := http.StatusOK // Change if different logic applies
	assert.Equal(t, expectedStatus, rec.Code)
}

func TestInsertMovie(t *testing.T) {
	movie := models.Movie{
		Title: "Sample Movie",
		Year:  2020,
	}
	body, err := json.Marshal(movie)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/api/movies/", bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(app.insert)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)  // Status might vary based on logic
}

func TestDeleteMovie(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies/{id}", app.delete)
	req, err := http.NewRequest("DELETE", "/api/movies/123", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	expectedStatus := http.StatusOK // This status may change based on the code analysis
	assert.Equal(t, expectedStatus, rec.Code)
}

