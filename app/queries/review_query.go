package queries

import (
	"context"
	"time"

	"github.com/Culture-Reviews-App/backend/app/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// ReviewQueries struct for queries related to the Review model.
type ReviewQueries struct {
	Collection *mongo.Collection
}

// CreateReview inserts a new review into the database.
func (q *ReviewQueries) CreateReview(review *models.Review) error {
	// Set context and timeout for the operation.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert the review into the collection.
	_, err := q.Collection.InsertOne(ctx, review)
	if err != nil {
		// Return error if the insertion fails.
		return err
	}

	// Return nil if insertion is successful.
	return nil
}

// ListReviews gets all reviews for the user.
func (q *ReviewQueries) ListReviews(userID uuid.UUID) ([]models.ReviewList, error) {
	// Set context and timeout for the operation.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define the filter to get reviews for the given userID
	filter := bson.M{"userid": userID}

	// Query the collection for reviews
	cursor, err := q.Collection.Find(ctx, filter)
	if err != nil {
		// Return error if the query fails.
		return nil, err
	}
	defer cursor.Close(ctx)

	// Define a slice to store the results
	var reviews []models.ReviewList

	// Iterate through the cursor and decode the documents
	if err = cursor.All(ctx, &reviews); err != nil {
		return nil, err
	}

	// Return the list of reviews
	return reviews, nil
}
