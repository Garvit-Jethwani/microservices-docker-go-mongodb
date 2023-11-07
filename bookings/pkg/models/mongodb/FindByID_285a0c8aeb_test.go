package mongodb

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mmorejon/microservices-docker-go-mongodb/bookings/pkg/models"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestFindByID_285a0c8aeb(t *testing.T) {
	testCases := []struct {
		name        string
		inputID     string
		expectedErr error
		mockFunc    func() (*BookingModel, *models.Booking)
	}{
		{
			name:        "Valid Input",
			inputID:     "614f4f064abf955bff540af0",
			expectedErr: nil,
			mockFunc: func() (*BookingModel, *models.Booking) {
				return &BookingModel,
					&models.Booking{
						ID:         primitive.ObjectIDFromHex("614f4f064abf955bff540af0"),
						UserID:     "user1",
						ShowtimeID: "showtime1",
						Movies:     []string{"movie1", "movie2"},
					}
			},
		},
		{
			name:        "Invalid Input",
			inputID:     "invalidID",
			expectedErr: mongo.ErrNoDocuments,
			mockFunc:    func() (*BookingModel, *models.Booking) { return nil, nil },
		},
		{
			name:        "Incorrect Input Format",
			inputID:     "12345",
			expectedErr: primitive.ErrInvalidHex,
			mockFunc:    func() (*BookingModel, *models.Booking) { return nil, nil },
		},
		{
			name:        "Database Unavailability",
			inputID:     "614f4f064abf955bff540af0",
			expectedErr: errors.New("database server unreachable"),
			mockFunc:    func() (*BookingModel, *models.Booking) { return nil, nil },
		},
		{
			name:        "Empty Input",
			inputID:     "",
			expectedErr: primitive.ErrInvalidHex,
			mockFunc:    func() (*BookingModel, *models.Booking) { return nil, nil },
		},
		{
			name:        "Null Document",
			inputID:     "614f4f064abf955bff540af1",
			expectedErr: nil,
			mockFunc: func() (*BookingModel, *models.Booking) {
				return &BookingModel,
					&models.Booking{
						ID: primitive.ObjectIDFromHex("614f4f064abf955bff540af1"),
					}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m, expectedBooking := tc.mockFunc()
			actualBooking, err := m.FindByID(tc.inputID)

			if err != nil && tc.expectedErr != nil {
				assert.ErrorIs(t, err, tc.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, expectedBooking, actualBooking)
				logrus.Info(fmt.Sprintf("Test case %s passed successfully", tc.name))
			}
		})
	}
}
