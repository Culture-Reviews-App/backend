package queries

import (
	"context"
	"time"

	"github.com/Culture-Reviews-App/backend/app/models"
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

	// Insert the user into the collection.
	_, err := q.Collection.InsertOne(ctx, review)
	if err != nil {
		// Return error if the insertion fails.
		return err
	}

	// Return nil if insertion is successful.
	return nil
}
