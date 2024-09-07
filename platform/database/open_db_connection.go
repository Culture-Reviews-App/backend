package database

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/Culture-Reviews-App/backend/app/queries"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Queries struct for collect all app queries.
type Queries struct {
	*queries.UserQueries // load queries from User model
}

// OpenDBConnection function for opening the MongoDB connection.
func OpenDBConnection(collectionName string) (*Queries, error) {
	// Get MongoDB connection details from environment variables.
	mongoURI := os.Getenv("MONGODB_URI")

	// Set client options.
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB.
	client, _ := mongo.Connect(clientOptions)

	// Create a context with timeout for connecting to the database.
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get the MongoDB collection.
	dbName := os.Getenv("MONGODB_DBNAME")
	collection := client.Database(dbName).Collection(collectionName)

	// Initialize the appropriate queries based on the collection name.
	switch collectionName {
	case "users":
		return &Queries{
			UserQueries: &queries.UserQueries{Collection: collection},
		}, nil
	default:
		// Handle the case where the collection name does not match any expected names.
		return nil, errors.New("invalid collection name")
	}
}
