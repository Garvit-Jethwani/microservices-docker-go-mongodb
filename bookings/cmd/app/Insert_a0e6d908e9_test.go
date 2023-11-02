// Test generated by RoostGPT for test go-mcvs using AI Type Azure Open AI and AI Model roost-gpt4-32k

/*
1. Test Scenario: Valid JSON
  - Input: Provide a valid JSON body that can successfully be decoded into the models.Booking struct.
  - Expectation: No error from 'json.NewDecoder().Decode()' function, successful completion of the 'insert' method, and a log entry indicating the successful creation of a new booking with the newly inserted booking ID.

2. Test Scenario: Invalid JSON
  - Input: Provide an invalid JSON body that cannot be decoded into the models.Booking struct.
  - Expectation: An error returned by 'json.NewDecoder().Decode()' function, the 'insert' method should handle the error, and generate a server error message.

3. Test Scenario: Insertion Failure
  - Input: Provide a valid JSON body that can be decoded into the models.Booking struct, but have the Insert method of the bookings object to fail deliberately.
  - Expectation: An error should be returned by the 'app.bookings.Insert(m)' function, the 'serverError' should handle this error, and a server error always be generated.

4. Test Scenario: Empty JSON
  - Input: Provide an empty JSON body.
  - Expectation: An error returned by 'json.NewDecoder().Decode()' function, the 'insert' method should be able to handle the error, and a server error message should be generated.

5. Test Scenario: Null Request Body
  - Input: Provide a null request body.
  - Expectation: An error returned by 'json.NewDecoder().Decode()' function, the 'insert' method should be able to handle the error, and a server error message should be generated.

6. Test Scenario: Successful Insertion
  - Input: Inject a mocked booking object that allows Insert method to execute successfully.
  - Expectation: A new booking should be created with a log entry displaying its ID. No error should be thrown.
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mmorejon/microservices-docker-go-mongodb/bookings/pkg/models"
)

func TestInsert_a0e6d908e9(t *testing.T) {
	tt := []struct {
		name        string
		body        interface{}
		insertErr   error
		expectedErr string
	}{
		{
			name: "Test Scenario: Valid JSON",
			body: models.Booking{ID: "123", RoomNo: 101,
				Price: float64(100), CheckIn: "2022-11-01",
				CheckOut: "2022-12-01", CustID: "222"},
		},
		{
			name: "Test Scenario: Invalid JSON",
			body: "invalid",
		},
		{
			name: "Test Scenario: Insertion Failure",
			body: models.Booking{ID: "123", RoomNo: 101,
				Price: float64(100), CheckIn: "2022-11-01",
				CheckOut: "2022-12-01", CustID: "222"},
			insertErr: fmt.Errorf("forced insert error"),
		},
		{
			name: "Test Scenario: Empty JSON",
		},
		{
			name: "Test Scenario: Null Request Body",
			body: nil,
		},
		{
			name: "Test Scenario: Successful Insertion",
			body: models.Booking{ID: "123", RoomNo: 101,
				Price: float64(100), CheckIn: "2022-11-01",
				CheckOut: "2022-12-01", CustID: "222"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tc.body)

			req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			rec := httptest.NewRecorder()
			r := mux.NewRouter()

			app := &application{
				bookings: &mockBookingModel{
					InsertFunc: func(m models.Booking) (*insertOneResult, error) { return nil, tc.insertErr },
				},
			}

			r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				app.insert(w, r)
			}).Methods("POST")
			r.ServeHTTP(rec, req)

			if tc.expectedErr != "" {
				if rec.Body.String() != tc.expectedErr {
					t.Errorf("expected error message %v got %v", tc.expectedErr, rec.Body.String())
				}
			} else {
				if rec.Code != http.StatusOK {
					t.Errorf("expected status %v got %v", http.StatusOK, rec.Code)
				}
			}
		})
	}
}
