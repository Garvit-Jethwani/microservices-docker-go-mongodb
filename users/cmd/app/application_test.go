// ********RoostGPT********
/*
Application Test generated by RoostGPT for test Application-Test-Golang using AI Type Open AI and AI Model gpt-4o


*/

// ********RoostGPT********
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// createTestApplication creates a test instance of the application
func createTestApplication() *application {
	return &application{
		errorLog: nil,
		infoLog:  nil,
		apis: apis{
			users:    "http://localhost:4000/api/users/",
			movies:   "http://localhost:4000/api/movies/",
			showtimes: "http://localhost:4000/api/showtimes/",
			bookings:  "http://localhost:4000/api/bookings/",
		},
	}
}

// createTestServer creates a test HTTP server
func createTestServer(app *application) *httptest.Server {
	return httptest.NewServer(app.routes())
}

func TestHomeHandler(t *testing.T) {
	app := createTestApplication()
	ts := createTestServer(app)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Fatalf("Could not send GET request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", resp.Status)
	}
}

func TestUsersListHandler(t *testing.T) {
	app := createTestApplication()
	ts := createTestServer(app)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users/list")
	if err != nil {
		t.Fatalf("Could not send GET request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", resp.Status)
	}
}

func TestUsersViewHandler(t *testing.T) {
	app := createTestApplication()
	ts := createTestServer(app)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users/view/1")
	if err != nil {
		t.Fatalf("Could not send GET request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", resp.Status)
	}
}

func TestMoviesListHandler(t *testing.T) {
	app := createTestApplication()
	ts := createTestServer(app)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/movies/list")
	if err != nil {
		t.Fatalf("Could not send GET request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", resp.Status)
	}
}

func TestMoviesViewHandler(t *testing.T) {
	app := createTestApplication()
	ts := createTestServer(app)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/movies/view/1")
	if err != nil {
		t.Fatalf("Could not send GET request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", resp.Status)
	}
}

func TestShowtimesListHandler(t *testing.T) {
	app := createTestApplication()
	ts := createTestServer(app)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/showtimes/list")
	if err != nil {
		t.Fatalf("Could not send GET request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", resp.Status)
	}
}

func TestShowtimesViewHandler(t *testing.T) {
	app := createTestApplication()
	ts := createTestServer(app)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/showtimes/view/1")
	if err != nil {
		t.Fatalf("Could not send GET request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", resp.Status)
	}
}

func TestBookingsListHandler(t *testing.T) {
	app := createTestApplication()
	ts := createTestServer(app)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/bookings/list")
	if err != nil {
		t.Fatalf("Could not send GET request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", resp.Status)
	}
}

func TestBookingsViewHandler(t *testing.T) {
	app := createTestApplication()
	ts := createTestServer(app)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/bookings/view/1")
	if err != nil {
		t.Fatalf("Could not send GET request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", resp.Status)
	}
}

