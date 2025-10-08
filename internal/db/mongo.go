package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Connect abre una conexión persistente con MongoDB 7.0
func Connect() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	opts := options.Client().ApplyURI(uri)
	c, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	if err := c.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("no se pudo conectar a MongoDB: %w", err)
	}

	fmt.Println("✅ Conectado a MongoDB 7.0 en", uri)
	client = c
	return c, nil
}

func GetCollection(name string) *mongo.Collection {
	dbName := os.Getenv("MONGO_DB")
	if dbName == "" {
		dbName = "goapi"
	}
	return client.Database(dbName).Collection(name)
}
