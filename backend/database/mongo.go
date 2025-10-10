package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"posttest/backend/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance *mongo.Client
	databaseName   string
	collectionName string
)

// Connect initializes a MongoDB client using env configuration and caches it for reuse.
func Connect() (*mongo.Client, error) {
	if clientInstance != nil {
		return clientInstance, nil
	}

	config.LoadEnv()

	uri := config.GetEnv("MONGODB_URI", "mongodb://localhost:27017")
	databaseName = config.GetEnv("MONGODB_DB", "geospatial")
	collectionName = config.GetEnv("MONGODB_COLLECTION", "polylines")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("mongo: unable to create client: %w", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("mongo: unable to connect: %w", err)
	}

	clientInstance = client
	log.Printf("mongo: connected to %s/%s", uri, databaseName)
	return clientInstance, nil
}

// Collection returns the polylines collection reference.
func Collection() (*mongo.Collection, error) {
	client, err := Connect()
	if err != nil {
		return nil, err
	}
	return client.Database(databaseName).Collection(collectionName), nil
}
