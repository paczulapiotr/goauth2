package database

import (
	"context"
	"time"

	"github.com/paczulapiotr/goauth2/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DefaultClient returns default mongo client
func DefaultClient() *mongo.Client {
	config := config.GetConfiguration()
	mongo, _ := CreateClient(config.Mongo)
	return mongo
}

// CreateClient opens mongodb connection and returns client
func CreateClient(connectionString string) (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)
	contx := context.TODO()

	// Connect to MongoDB
	client, err := mongo.Connect(contx, clientOptions)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func createContext() *context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	return &ctx
}
