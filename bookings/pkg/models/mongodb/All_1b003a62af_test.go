package mongodb

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"
	"log"

	"github.com/mmorejon/microservices-docker-go-mongodb/bookings/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const connectionString = "mongodb://localhost:27017"

func TestAll_1b003a62af(t *testing.T) {

	client, _ := mongo.NewClient(options.Client().ApplyURI(connectionString))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := client.Connect(ctx)

	fmt.Fprintf(os.Stdout, "Connecting to MongoDB...")
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not connect to MongoDB: %v\n", err)
		return
	}
	fmt.Fprintf(os.Stdout, "Connected to MongoDB\n")

	coll := client.Database("test").Collection("events")

	tests := []struct {
		name       string
		want       []models.Booking
		insertData bool
		wantErr    bool
		errMsg     string
	}{
		{
			"Test scenario 1: Fetch all bookings",
			[]models.Booking{{ID: primitive.ObjectID{}, Name: "TestBooking1"},
				{ID: primitive.ObjectID{}, Name: "TestBooking2"}},
			true,
			false,
			"",
		},
		{
			"Test scenario 2: Empty collection",
			[]models.Booking{},
			false,
			false,
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.insertData {
				_, _ = coll.InsertMany(ctx, bson.A{
					bson.M{"name": "TestBooking1"},
					bson.M{"name": "TestBooking2"},
				})
			}
			b := BookingModel{C: coll}
			got, err := b.All()
			if (err != nil) != tt.wantErr {
				t.Errorf("BookingModel.All() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.errMsg != err.Error() {
				t.Errorf("BookingModel.All() errMsg = %v, wantErrMsg %v", err, tt.wantErr)
				return
			}
			// Other assertions can be added here to validate the output
			if got, want := len(got), len(tt.want); got != want {
				t.Errorf("expected %d bookings, got %d", want, got)
			} else {
                log.Println("Test success: expected and got ", want)
            }
			// Test clear data
			_, _ = coll.DeleteMany(ctx, bson.M{})
		})
	}
}
