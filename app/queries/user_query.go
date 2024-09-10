package queries

import (
	"context"
	"time"

	"github.com/Culture-Reviews-App/backend/app/models" // Adjust the import path as necessary
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// UserQueries struct for queries related to the User model.
type UserQueries struct {
	Collection *mongo.Collection
}

// GetUserByID retrieves a user by their ID.
func (q *UserQueries) GetUserByID(id uuid.UUID) (*models.User, error) {
	// Define a variable to hold the result.
	var user models.User

	// Define a filter to find the user by ID.
	filter := bson.M{"id": id}

	// Execute the query.
	err := q.Collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		// Return nil and error if user is not found or other errors occur.
		return nil, err
	}

	// Return the user and nil as error.
	return &user, nil
}

// GetUserByEmail retrieves a user by their email.
func (q *UserQueries) GetUserByEmail(email string) (*models.User, error) {
	// Define a variable to hold the result.
	var user models.User

	// Define a filter to find the user by email.
	filter := bson.M{"email": email}

	// Execute the query.
	err := q.Collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		// Return nil and error if user is not found or other errors occur.
		return nil, err
	}

	// Return the user and nil as error.
	return &user, nil
}

// CreateUser inserts a new user into the database.
func (q *UserQueries) CreateUser(user *models.User) error {
	// Set context and timeout for the operation.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert the user into the collection.
	_, err := q.Collection.InsertOne(ctx, user)
	if err != nil {
		// Return error if the insertion fails.
		return err
	}

	// Return nil if insertion is successful.
	return nil
}

// GetUserByEmail retrieves a user by their email.
func (q *UserQueries) GetUserByUsername(username string) (*models.User, error) {
	// Define a variable to hold the result.
	var user models.User

	// Define a filter to find the user by email.
	filter := bson.M{"username": username}

	// Execute the query.
	err := q.Collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		// Return nil and error if user is not found or other errors occur.
		return nil, err
	}

	// Return the user and nil as error.
	return &user, nil
}
