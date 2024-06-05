// ********RoostGPT********
/*
Test generated by RoostGPT for test go-with-application-test using AI Type Open AI and AI Model gpt-4o

ROOST_METHOD_HASH=All_4c4a1c9150
ROOST_METHOD_SIG_HASH=All_1b003a62af

The function `All` is defined in the `mongodb` package and retrieves all records from the bookings collection in a MongoDB database. Here is the `All` function and related context:

### Package and Imports

```go
package mongodb

import (
	"context"
	"errors"

	"github.com/mmorejon/microservices-docker-go-mongodb/bookings/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
```

### Struct Definition

```go
// BookingModel represents a MongoDB collection with a booking data model
type BookingModel struct {
	C *mongo.Collection
}
```

### Function All

```go
// All method will be used to get all records from bookings table
func (m *BookingModel) All() ([]models.Booking, error) {
	// Define variables
	ctx := context.TODO()
	b := []models.Booking{}

	// Find all bookings
	bookingCursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = bookingCursor.All(ctx, &b)
	if err != nil {
		return nil, err
	}

	return b, err
}
```

Based on the function definition, the file package, imports, and struct definition, I’ll craft several test scenarios:

### Test Scenarios for the `All` Function

Scenario 1: Fetch All Bookings When Collection is Empty

- **Description**: Tests if the `All` function handles an empty collection correctly.
- **Execution**:
  - **Arrange**: Mock the MongoDB collection to return an empty cursor.
  - **Act**: Call the `All` method.
  - **Assert**: Verify that the returned slice is empty and no errors are returned.
- **Validation**:
  - **Justify**: Ensures the function handles the edge case of no records gracefully.
  - **Importance**: Confirms the system's stability with no data.

Scenario 2: Fetch Multiple Bookings

- **Description**: Ensure the `All` function correctly retrieves all bookings from the collection.
- **Execution**:
  - **Arrange**: Mock the MongoDB collection to return a cursor with multiple booking documents.
  - **Act**: Call the `All` function.
  - **Assert**: Verify that the returned slice contains the correct number of bookings with accurate data.
- **Validation**:
  - **Justify**: Validates that the function correctly retrieves and unmarshals multiple records.
  - **Importance**: Ensures the function's correctness in common use cases.

Scenario 3: Handle Database Connection Error

- **Description**: Simulates a database connection error when fetching bookings.
- **Execution**:
  - **Arrange**: Mock the MongoDB collection to return a connection error.
  - **Act**: Call the `All` method.
  - **Assert**: Verify that the function returns an error and no bookings.
- **Validation**:
  - **Justify**: Tests the function's robustness in handling database connectivity issues.
  - **Importance**: Ensures the system's reliability and error reporting.

Scenario 4: Handle Cursor Iteration Error

- **Description**: Simulates an error during the cursor iteration when fetching bookings.
- **Execution**:
  - **Arrange**: Mock the MongoDB cursor to return an iteration error.
  - **Act**: Call the `All` function.
  - **Assert**: Verify that the function returns an iteration error and no bookings.
- **Validation**:
  - **Justify**: Ensures the function can handle and propagate cursor errors.
  - **Importance**: Maintains reliability and proper error propagation.

### Example Test Code

Here’s a sample test structure in Go for some of these scenarios:

```go
package mongodb_test

import (
	"context"
	"testing"
	"github.com/mmorejon/microservices-docker-go-mongodb/bookings/pkg/models"
	"github.com/mmorejon/microservices-docker-go-mongodb/bookings/pkg/models/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"github.com/stretchr/testify/assert"
)

func TestAll_EmptyCollection(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("empty collection", func(mt *mtest.T) {
		// Arrange
		collection := mt.Coll
		model := &mongodb.BookingModel{C: collection}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "mocked.bookings", mtest.FirstBatch))

		// Act
		bookings, err := model.All()

		// Assert
		assert.NoError(t, err)
		assert.Empty(t, bookings)
	})
}

func TestAll_MultipleBookings(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("multiple bookings", func(mt *mtest.T) {
		// Arrange
		collection := mt.Coll
		model := &mongodb.BookingModel{C: collection}

		bookings := []models.Booking{
			{ID: "1", Name: "Booking 1"},
			{ID: "2", Name: "Booking 2"},
		}

		firstBatch := mtest.CreateCursorResponse(1, "mocked.bookings", mtest.FirstBatch, bookings[0])
		secondBatch := mtest.CreateCursorResponse(1, "mocked.bookings", mtest.NextBatch, bookings[1])
		mt.AddMockResponses(firstBatch, secondBatch)

		// Act
		result, err := model.All()

		// Assert
		assert.NoError(t, err)
		assert.Len(t, result, 2)
	})
}

func TestAll_ConnectionError(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("connection error", func(mt *mtest.T) {
		// Arrange
		collection := mt.Coll
		model := &mongodb.BookingModel{C: collection}

		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "connection error",
		}))

		// Act
		bookings, err := model.All()

		// Assert
		assert.Error(t, err)
		assert.Nil(t, bookings)
	})
}
```

These tests cover key cases including an empty collection, multiple records, and connection errors to ensure the `All` function operates properly under various conditions.
*/

// ********RoostGPT********
package mongodb_test

import (
	"testing"
	"github.com/mmorejon/microservices-docker-go-mongodb/bookings/pkg/models"
	"github.com/mmorejon/microservices-docker-go-mongodb/bookings/pkg/models/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	tests := []struct {
		name           string
		setupMock      func(mt *mtest.T)
		expectedErr    error
		expectedResult []models.Booking
	}{
		{
			name: "empty collection",
			setupMock: func(mt *mtest.T) {
				// Arrange
				mt.AddMockResponses(mtest.CreateCursorResponse(1, "mocked.bookings", mtest.FirstBatch))
			},
			expectedErr:    nil,
			expectedResult: []models.Booking{},
		},
		{
			name: "multiple bookings",
			setupMock: func(mt *mtest.T) {
				// Arrange
				bookings := []models.Booking{
					{ID: primitive.NewObjectID(), UserID: "user1", ShowtimeID: "showtime1", Movies: []string{"movie1", "movie2"}},
					{ID: primitive.NewObjectID(), UserID: "user2", ShowtimeID: "showtime2", Movies: []string{"movie3", "movie4"}},
				}
				mt.AddMockResponses(
					mtest.CreateCursorResponse(0, "mocked.bookings", mtest.FirstBatch, bson.D{
						{Key: "_id", Value: bookings[0].ID},
						{Key: "userid", Value: bookings[0].UserID},
						{Key: "showtimeid", Value: bookings[0].ShowtimeID},
						{Key: "movies", Value: bookings[0].Movies},
					}),
					mtest.CreateCursorResponse(0, "mocked.bookings", mtest.NextBatch, bson.D{
						{Key: "_id", Value: bookings[1].ID},
						{Key: "userid", Value: bookings[1].UserID},
						{Key: "showtimeid", Value: bookings[1].ShowtimeID},
						{Key: "movies", Value: bookings[1].Movies},
					}),
				)
			},
			expectedErr: nil,
			expectedResult: []models.Booking{
				{ID: primitive.NewObjectID(), UserID: "user1", ShowtimeID: "showtime1", Movies: []string{"movie1", "movie2"}},
				{ID: primitive.NewObjectID(), UserID: "user2", ShowtimeID: "showtime2", Movies: []string{"movie3", "movie4"}},
			},
		},
		{
			name: "connection error",
			setupMock: func(mt *mtest.T) {
				// Arrange
				mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
					Code:    11000,
					Message: "connection error",
				}))
			},
			expectedErr:    mongo.CommandError{Code: 11000, Message: "connection error"},
			expectedResult: nil,
		},
		{
			name: "cursor iteration error",
			setupMock: func(mt *mtest.T) {
				// Arrange
				bookings := []models.Booking{
					{ID: primitive.NewObjectID(), UserID: "user1", ShowtimeID: "showtime1", Movies: []string{"movie1", "movie2"}},
				}
				firstBatch := mtest.CreateCursorResponse(1, "mocked.bookings", mtest.FirstBatch, bson.D{
					{Key: "_id", Value: bookings[0].ID},
					{Key: "userid", Value: bookings[0].UserID},
					{Key: "showtimeid", Value: bookings[0].ShowtimeID},
					{Key: "movies", Value: bookings[0].Movies},
				})
				mt.AddMockResponses(firstBatch)
				mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
					Code:    2,
					Message: "cursor iteration error",
				}))
			},
			expectedErr:    mongo.CommandError{Code: 2, Message: "cursor iteration error"},
			expectedResult: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
			defer mt.Close()
			collection := mt.Coll
			model := &mongodb.BookingModel{C: collection}

			tt.setupMock(mt)

			// Act
			result, err := model.All()

			// Assert
			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}
		})
	}
}

