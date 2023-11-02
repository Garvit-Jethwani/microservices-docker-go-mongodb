// Test generated by RoostGPT for test go-mcvs using AI Type Azure Open AI and AI Model roost-gpt4-32k

/*
1. Test case when `id` is valid:
  - Scenario: The client sends a valid `id` which exists in the database.
  - Expected: The bookings record with the specified `id` is deleted successfully. The log prints an appropriate message indicating the number of bookings deleted.

2. Test case when `id` is not valid:
  - Scenario: The client sends an `id` that is incorrectly formatted or not valid.
  - Expected: An error should be returned and the server error should be logged.

3. Test case when `id` does not exist:
  - Scenario: The client sends an `id` that is well-formatted, but does not exist in the database.
  - Expected: The method should handle this correctly, potentially by returning a "Not Found" message or by simply logging that 0 bookings were deleted.

4. Test case for deleting multiple bookings:
  - Scenario: The client sends a request to delete multiple bookings with the provided `id`.
  - Expected: All bookings with the matched `id` should be deleted and the logs should indicate the correct number of bookings deleted.

5. Test scenario when database connection is lost:
  - Scenario: During the execution of the delete method, the database connection is lost.
  - Expected: An error should be returned and the app should log a server error message.
*/
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

// Mocking the bookings
type mockBookings struct{}

func (m *mockBookings) Delete(id string) (string, error) {
	if id == "invalid_id" {
		return "", fmt.Errorf("invalid id")
	}

	if id == "not_exist_id" {
		return "0", nil
	}

	if id == "multiple_id" {
		return "3", nil
	}

	if id == "database_error" {
		return "", fmt.Errorf("database connection lost")
	}

	return "1", nil
}

func TestDelete_b339da421f(t *testing.T) {
	// Define test cases
	tests := []struct {
		name       string
		id         string
		wantLog    string
		wantStatus int
	}{
		{"Valid ID", "123", "Have been eliminated 1 booking(s)", http.StatusOK},
		{"Invalid ID", "invalid_id", "ERROR : invalid id", http.StatusInternalServerError},
		{"Not Exist ID", "not_exist_id", "Have been eliminated 0 booking(s)", http.StatusOK},
		{"Delete Multiple ID", "multiple_id", "Have been eliminated 3 booking(s)", http.StatusOK},
		{"Database Error", "database_error", "ERROR : database connection lost", http.StatusInternalServerError},
	}

	app := &application{
		errorLog: log.New(ioutil.Discard, "", 0),
		infoLog:  log.New(ioutil.Discard, "", 0),
		bookings: &mockBookings{},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/delete/%s", tt.id), nil)
			responseRecorder := httptest.NewRecorder()

			// Add the id to the URL path
			request = mux.SetURLVars(request, map[string]string{"id": tt.id})

			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// run the delete method
			app.delete(responseRecorder, request)

			// closing and resetting os.Stdout
			w.Close()
			out, _ := ioutil.ReadAll(r)
			os.Stdout = oldStdout

			// check for the http response status
			if responseRecorder.Code != tt.wantStatus {
				t.Errorf("want %d; got %d", tt.wantStatus, responseRecorder.Code)
			}

			// check for the log messages
			if !bytes.Contains(out, []byte(tt.wantLog)) {
				t.Errorf("expected log message to include %q", tt.wantLog)
			}
		})
	}
}
