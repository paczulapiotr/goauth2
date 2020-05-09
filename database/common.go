package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func closeConnection(mongo *mongo.Client, ctx *context.Context) error {
	return mongo.Disconnect(*ctx)
}
